package midiplayer

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ncirocco/midino/midimessages"
	"github.com/ncirocco/midino/midiparser"
	"github.com/ncirocco/midino/miditiming"
	"github.com/rakyll/portmidi"
)

const midiChannels int64 = 16

// PlayMIDI plays the given MIDI
func PlayMIDI(midi *midiparser.Midi) {
	portmidi.Initialize()
	out, err := portmidi.NewOutputStream(2, 1024, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	defer portmidi.Terminate()

	// In case that the user interrupts the program (ctrl+c)
	// this will stop properly all the playing notes.
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		StopAllNotes(out)
		os.Exit(1)
	}()

	play := make(map[int64][]*midiparser.Event)
	m := int64(0)
	for _, track := range midi.Tracks {
		t := int64(0)
		for _, event := range track.Events {
			t = t + event.DeltaTicks
			event.AbsoluteTicks = t
			play[t] = append(play[t], event)
		}
		if m < t {
			m = t
		}

	}

	lastEventTicks := int64(0)
	msPerTick := miditiming.TempoAndPpqnToMS(midi.Tempo, midi.Ppqn)

	for i := int64(0); i < m; i++ {
		for _, event := range play[i] {
			deltaTicks := event.AbsoluteTicks - lastEventTicks
			lastEventTicks = event.AbsoluteTicks
			time.Sleep(time.Duration(msPerTick * deltaTicks))
			out.WriteShort(
				event.Status+event.Channel,
				event.FirstDataByte,
				event.SecondDataByte,
			)
		}
	}

	StopAllNotes(out)
}

// StopAllNotes stops all the notes in all the channels inmidiatelly
func StopAllNotes(out *portmidi.Stream) {
	for channel := int64(0); channel < midiChannels; channel++ {
		m := midimessages.StopAll(channel)
		out.WriteShort(
			m.StatusByte,
			m.DataByte1,
			m.DataByte2,
		)
	}
}
