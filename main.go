package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/algoGuy/EasyMIDI/smf"
	"github.com/algoGuy/EasyMIDI/smfio"
	"github.com/rakyll/portmidi"
)

const defaultTempo int64 = 500000

const tempoMetaType uint8 = 0x51

type midiSong struct {
	tracks map[int64]*track
	bpm    int64
	ppq    int64
}

type track struct {
	midiEvents []*midiEvent
}

type midiEvent struct {
	status         int64
	channel        int64
	firstDataByte  int64
	secondDataByte int64
	deltaTicks     int64
}

func main() {
	midiFile := os.Args[1]
	file, err := os.Open(midiFile)
	if err != nil {
		fmt.Println("File is invalid")

		os.Exit(1)
	}
	defer file.Close()

	midi, err := smfio.Read(bufio.NewReader(file))
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	div := midi.GetDivision()

	ms := midiSong{
		tracks: make(map[int64]*track),
		ppq:    int64(div.GetTicks()),
		bpm:    calculateBpm(defaultTempo),
	}

	for i := int64(0); i < int64(midi.GetTracksNum()); i++ {
		t := midi.GetTrack(uint16(i))
		ms.tracks[i] = &track{}

		events := t.GetAllEvents()
		for _, e := range events {
			m, ok := e.(*smf.MIDIEvent)
			if ok {
				ms.tracks[i].midiEvents = append(ms.tracks[i].midiEvents, handleMidiEvent(m))

				continue
			}
			metaEvent, ok := e.(*smf.MetaEvent)
			if ok {
				if metaEvent.GetMetaType() == tempoMetaType {
					tempo, _ := strconv.ParseInt(hex.EncodeToString(metaEvent.GetData()), 16, 64)
					ms.bpm = calculateBpm(tempo)
				}
			}
		}
	}

	playMIDI(&ms)
}

func handleMidiEvent(event *smf.MIDIEvent) *midiEvent {
	firstDataByte := int64(event.GetData()[0])
	secondDataByte := int64(0)
	if len(event.GetData()) > 1 {
		secondDataByte = int64(event.GetData()[1])
	}

	return &midiEvent{
		status:         int64(event.GetStatus()),
		channel:        int64(event.GetChannel()),
		firstDataByte:  firstDataByte,
		secondDataByte: secondDataByte,
		deltaTicks:     int64(event.GetDTime()),
	}
}

func calculateBpm(tempo int64) int64 {
	return (60 * 1000000) / tempo
}

func ticksToSeconds(deltaTime int64, bpm int64, ppq int64) int64 {
	return deltaTime * 60000 / (bpm * ppq)
}

func playMIDI(midi *midiSong) {
	portmidi.Initialize()
	out, err := portmidi.NewOutputStream(2, 1024, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	defer portmidi.Terminate()

	play := make(map[int64][]*midiEvent)
	max := int64(0)

	for _, track := range midi.tracks {
		t := int64(0)
		for _, event := range track.midiEvents {
			t = t + event.deltaTicks
			play[t] = append(play[t], event)
		}

		fmt.Println(t)
		if max < t {
			max = t
		}
	}

	s := time.Duration(ticksToSeconds(1, midi.bpm, midi.ppq)) * time.Millisecond
	for i := int64(0); i < max; i++ {
		time.Sleep(s)
		if len(play[i]) > 0 {
			for _, event := range play[i] {
				out.WriteShort(
					event.status+event.channel,
					event.firstDataByte,
					event.secondDataByte,
				)
			}
		}
	}
}
