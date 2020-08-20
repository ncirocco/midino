package midiparser

import (
	"encoding/hex"
	"strconv"

	"github.com/algoGuy/EasyMIDI/smf"
	"github.com/ncirocco/midi-player/miditiming"
)

const tempoMetaType uint8 = 0x51

// Midi holds all the information for a MIDI file
type Midi struct {
	Tracks map[int]*Track
	Bpm    int64
	Ppqn   int64
	Tempo  int64
}

// Track holds all the events for a given track in a MIDI file
type Track struct {
	Events []*Event
}

// Event holds information such as playing a note or adjusting a MIDI channel's modulation value
type Event struct {
	Status         int64
	Channel        int64
	FirstDataByte  int64
	SecondDataByte int64
	DeltaTicks     int64
	AbsoluteTicks  int64
}

// ParseMidiFile parses the file and returns a Midi struct
func ParseMidiFile(midiFile *smf.MIDIFile) *Midi {
	div := midiFile.GetDivision()
	ms := Midi{
		Tracks: make(map[int]*Track),
		Ppqn:   int64(div.GetTicks()),
		Bpm:    miditiming.CalculateBpm(miditiming.DefaultTempo),
		Tempo:  miditiming.DefaultTempo,
	}

	for i := 0; i < int(midiFile.GetTracksNum()); i++ {
		t := midiFile.GetTrack(uint16(i))
		ms.Tracks[i] = &Track{}

		events := t.GetAllEvents()
		for _, e := range events {
			m, ok := e.(*smf.MIDIEvent)
			if ok {
				ms.Tracks[i].Events = append(ms.Tracks[i].Events, handleMidiEvent(m))

				continue
			}
			metaEvent, ok := e.(*smf.MetaEvent)
			if ok {
				if metaEvent.GetMetaType() == tempoMetaType {
					tempo, _ := strconv.ParseInt(hex.EncodeToString(metaEvent.GetData()), 16, 64)
					ms.Bpm = miditiming.CalculateBpm(tempo)
					ms.Tempo = tempo
				}
			}
		}
	}

	return &ms
}

func handleMidiEvent(event *smf.MIDIEvent) *Event {
	firstDataByte := int64(event.GetData()[0])
	secondDataByte := int64(0)
	if len(event.GetData()) > 1 {
		secondDataByte = int64(event.GetData()[1])
	}

	return &Event{
		Status:         int64(event.GetStatus()),
		Channel:        int64(event.GetChannel()),
		FirstDataByte:  firstDataByte,
		SecondDataByte: secondDataByte,
		DeltaTicks:     int64(event.GetDTime()),
	}
}
