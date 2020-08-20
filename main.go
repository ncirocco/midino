package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ncirocco/midino/midiparser"
	"github.com/ncirocco/midino/midiplayer"

	"github.com/algoGuy/EasyMIDI/smfio"
)

func main() {
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("File is invalid")

		os.Exit(1)
	}
	defer file.Close()

	midiFile, err := smfio.Read(bufio.NewReader(file))
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	midi := midiparser.ParseMidiFile(midiFile)
	midiplayer.PlayMIDI(midi)
}
