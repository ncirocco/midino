package midiplayer

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/ncirocco/midino/midimessages"
	"github.com/ncirocco/midino/midiparser"
	"github.com/ncirocco/midino/miditiming"
	"github.com/rakyll/portmidi"
)

const midiChannels int64 = 16

// Player contains all the necessary logic to play a MIDI song
type Player struct {
	currentSong       int
	currentSongLength int64
	currentSongTime   int64
	width             int
	height            int
	pause             bool
	event             chan int
	out               *portmidi.Stream
	playlist          []string
	playlistUI        *widgets.List
}

// NewPlayer returns an instance of a Player
func NewPlayer(deviceID int64) (*Player, error) {
	portmidi.Initialize()
	out, err := portmidi.NewOutputStream(portmidi.DeviceID(deviceID), 1024, 0)
	if err != nil {
		return nil, err
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	event := make(chan int)
	width, height := ui.TerminalDimensions()

	l := widgets.NewList()
	l.Title = "Playlist"
	l.Rows = []string{}
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false

	p := &Player{
		width:      width,
		height:     height,
		event:      event,
		out:        out,
		playlistUI: l,
	}

	go p.hotkeys()

	return p, nil
}

// Play starts the player
func (p *Player) Play() {
	for {
		midi, err := midiparser.ParseMidiFile(p.playlist[p.currentSong])
		if err != nil {
			log.Fatal(err)
		}
		if !p.playMIDI(midi) {
			break
		}
	}
}

// playMIDI returns false if it should not play a next song
func (p *Player) playMIDI(midi *midiparser.Midi) bool {
	// reset all channels to the default value to avoid problems
	// with files that assume all values are already in its default
	resetChannels(p.out)
	p.updateScreen()

	play := make(map[int64][]*midiparser.Event)
	totalTicks := int64(0)
	for _, track := range midi.Tracks {
		t := int64(0)
		for _, event := range track.Events {
			t = t + event.DeltaTicks
			event.AbsoluteTicks = t
			play[t] = append(play[t], event)
		}
		if totalTicks < t {
			totalTicks = t
		}
	}

	lastEventTicks := int64(0)
	nsPerTick := miditiming.TempoAndPpqnToNS(midi.Tempo, midi.Ppqn)

	fq := 1000000000 / nsPerTick
	p.currentSongLength = totalTicks * nsPerTick

	for i := int64(0); i < totalTicks; i++ {
		for p.pause {
			stopAllNotes(p.out)
			time.Sleep(100 * time.Millisecond)
		}

		if i%fq == 0 {
			p.currentSongTime = i * nsPerTick
			p.updateScreen()
		}

		select {
		case msg := <-p.event:
			switch msg {
			case 0:
				stopAllNotes(p.out)
				return false
			case 1:
				stopAllNotes(p.out)
				return true
			}
		default:
		}

		for _, event := range play[i] {
			deltaTicks := event.AbsoluteTicks - lastEventTicks
			lastEventTicks = event.AbsoluteTicks
			time.Sleep(time.Duration(nsPerTick * deltaTicks))
			p.out.WriteShort(
				event.Status+event.Channel,
				event.FirstDataByte,
				event.SecondDataByte,
			)
		}
	}

	// mutes all notes that are playing that don't have a mute note message
	muteAllNotes(p.out)
	// give some time in between songs
	time.Sleep(time.Duration(2 * time.Second))
	// just in case force stop any remaining sounds from the previous song
	stopAllNotes(p.out)

	if p.currentSong < len(p.playlist)-1 {
		p.currentSong++
	} else {
		p.currentSong = 0
	}

	return true
}

// AddToPlaylist Adds the path of a MIDI file to the playlist
func (p *Player) AddToPlaylist(path string) {
	p.playlist = append(p.playlist, path)
	p.playlistUI.Rows = append(p.playlistUI.Rows, filepath.Base(path))
}

// ListDevices lists all the available output MIDI devices and their IDs
func ListDevices() error {
	err := portmidi.Initialize()
	if err != nil {
		return err
	}

	defer portmidi.Terminate()

	numDevices := portmidi.CountDevices()

	fmt.Println("Available output MIDI devices")
	for i := 0; i < numDevices; i++ {
		info := portmidi.Info(portmidi.DeviceID(i))

		if info.IsOutputAvailable {
			fmt.Println("ID:", i, "Name:", info.Name)
		}
	}

	return nil
}

func (p *Player) updateScreen() {
	currentSong := widgets.NewParagraph()
	currentSong.Text = p.playlistUI.Rows[p.currentSong] +
		" - " +
		secondsToMinutes(p.currentSongTime/int64(time.Second)) +
		" - " +
		secondsToMinutes(p.currentSongLength/int64(time.Second))

	currentSong.SetRect(0, 0, p.width, 3)

	ui.Render(currentSong)

	hotkeys := widgets.NewParagraph()
	hotkeys.Text = "MIDINO - " +
		"q - quit | " +
		"l - next song | " +
		"h - prev song | " +
		"space - pause/resume"

	hotkeys.SetRect(0, p.height, p.width, p.height-3)

	ui.Render(hotkeys)

	p.playlistUI.SetRect(0, 3, p.width, p.height-3)
	ui.Render(p.playlistUI)
}

func (p *Player) hotkeys() {
	previousKey := ""
	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			switch e.ID {
			case "<Space>":
				p.pause = !p.pause
			case "j", "<Down>":
				p.playlistUI.ScrollDown()
			case "k", "<Up>":
				p.playlistUI.ScrollUp()
			case "<C-d>":
				p.playlistUI.ScrollHalfPageDown()
			case "<C-u>":
				p.playlistUI.ScrollHalfPageUp()
			case "<C-f>", "<PageDown>":
				p.playlistUI.ScrollPageDown()
			case "<C-b>", "<PageUp>":
				p.playlistUI.ScrollPageUp()
			case "<Home>":
				p.playlistUI.ScrollTop()
			case "G", "<End>":
				p.playlistUI.ScrollBottom()
			case "g":
				if previousKey == "g" {
					p.playlistUI.ScrollTop()
				}
			case "h":
				if p.currentSong > 0 {
					p.currentSong--
				}

				p.event <- 1
			case "l":
				if p.currentSong < len(p.playlist)-1 {
					p.currentSong++
				}
				p.event <- 1
			case "<C-c>", "q":
				p.event <- 0
			case "<Enter>":
				p.currentSong = p.playlistUI.SelectedRow

				p.event <- 1
			}

			if previousKey == "g" {
				previousKey = ""
			} else {
				previousKey = e.ID
			}
		}
		if e.ID == "<Resize>" {
			payload := e.Payload.(ui.Resize)
			width, height := payload.Width, payload.Height
			p.width = width
			p.height = height
		}
		p.updateScreen()
	}
}

// Close stops all notesm closes the output stream and terminates portmidi.
func (p *Player) Close() {
	stopAllNotes(p.out)
	p.out.Close()
	portmidi.Terminate()
	ui.Close()
}

func secondsToMinutes(inSeconds int64) string {
	minutes := inSeconds / 60
	seconds := inSeconds % 60
	str := fmt.Sprintf("%02d:%02d", minutes, seconds)
	return str
}

// resetChannels resets all channels to the default instrument and clears all controllers
func resetChannels(out *portmidi.Stream) {
	for channel := int64(0); channel < midiChannels; channel++ {
		m := midimessages.SelectInstrument(channel, midimessages.AcousticGrandPiano)
		out.WriteShort(
			m.StatusByte,
			m.DataByte1,
			m.DataByte2,
		)
		m = midimessages.ClearAllControllersForChannel(channel)
		out.WriteShort(
			m.StatusByte,
			m.DataByte1,
			m.DataByte2,
		)
	}
}

func stopAllNotes(out *portmidi.Stream) {
	for channel := int64(0); channel < midiChannels; channel++ {
		m := midimessages.StopAll(channel)
		out.WriteShort(
			m.StatusByte,
			m.DataByte1,
			m.DataByte2,
		)
	}
}

func muteAllNotes(out *portmidi.Stream) {
	for channel := int64(0); channel < midiChannels; channel++ {
		m := midimessages.MuteAll(channel)
		out.WriteShort(
			m.StatusByte,
			m.DataByte1,
			m.DataByte2,
		)
	}
}
