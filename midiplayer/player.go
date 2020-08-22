package midiplayer

import (
	"fmt"
	"time"

	"github.com/ncirocco/midino/midimessages"
	"github.com/ncirocco/midino/midiparser"
	"github.com/ncirocco/midino/miditiming"
	"github.com/rakyll/portmidi"
)

const midiChannels int64 = 16

// Player contains all the necessary logic to play a MIDI song
type Player struct {
	out *portmidi.Stream
}

// NewPlayer returns an instance of a Player
func NewPlayer() (*Player, error) {
	portmidi.Initialize()
	out, err := portmidi.NewOutputStream(2, 1024, 0)
	if err != nil {
		return nil, err
	}

	return &Player{
		out: out,
	}, nil
}

// PlayMIDI plays the given MIDI
func (p *Player) PlayMIDI(midi *midiparser.Midi) {
	// reset all channels to the default value to avoid problems
	// with files that assume all values are already in its default
	resetChannels(p.out)

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

	displayInfo(midi, totalTicks*nsPerTick)

	for i := int64(0); i < totalTicks; i++ {
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
}

func displayInfo(midi *midiparser.Midi, duration int64) {
	fmt.Println(
		midi.Name,
		"-",
		secondsToMinutes(duration/int64(time.Second)),
	)
}

func secondsToMinutes(inSeconds int64) string {
	minutes := inSeconds / 60
	seconds := inSeconds % 60
	str := fmt.Sprintf("%02d:%02d", minutes, seconds)
	return str
}

// Close stops all notesm closes the output stream and terminates portmidi.
func (p *Player) Close() {
	stopAllNotes(p.out)
	p.out.Close()
	portmidi.Terminate()
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
