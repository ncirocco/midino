package main

import (
	"bufio"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/ncirocco/midino/midiplayer"
)

func main() {
	if len(os.Args[1]) == 2 && os.Args[1] == "-l" {
		err := midiplayer.ListDevices()
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	if len(os.Args) != 3 {
		log.Fatal("Missing DeviceID and file or playlist argument.")
	}

	deviceID, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	p, err := midiplayer.NewPlayer(int64(deviceID))
	if err != nil {
		log.Fatal("Invalid deviceID. ")
	}
	defer p.Close()

	registerInterrupSignal(p)

	if strings.HasSuffix(strings.ToLower(os.Args[2]), ".mid") {
		p.AddToPlaylist(os.Args[2])
	} else {
		err = playPlaylist(os.Args[2], p)
		if err != nil {
			log.Fatal(err)
		}
	}

	p.Play()
}

func playPlaylist(path string, p *midiplayer.Player) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p.AddToPlaylist(scanner.Text())
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
