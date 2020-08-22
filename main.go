package main

import (
	"bufio"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ncirocco/midino/midiparser"
	"github.com/ncirocco/midino/midiplayer"
)

func main() {
	if len(os.Args) == 0 {
		log.Fatal("Missing file or playlist argument.")
	}

	p, err := midiplayer.NewPlayer()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	registerInterrupSignal(p)

	if strings.HasSuffix(strings.ToLower(os.Args[1]), ".mid") {
		midi, err := midiparser.ParseMidiFile(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		p.PlayMIDI(midi)

		return
	}

	err = playPlaylist(os.Args[1], p)
	if err != nil {
		log.Fatal(err)
	}
}

func playPlaylist(path string, p *midiplayer.Player) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		midi, err := midiparser.ParseMidiFile(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		p.PlayMIDI(midi)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// registerInterrupSignal registers what will happen if the program is interrupted
func registerInterrupSignal(p *midiplayer.Player) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		p.Close()
		os.Exit(1)
	}()
}
