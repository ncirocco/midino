# MIDINO player

**MIDINO** is a Linux CLI MIDI player for MIDI sound modules, like the [Roland Sound Canvas](https://en.wikipedia.org/wiki/Roland_Sound_Canvas) and the [Yamaha MU-series](https://en.wikipedia.org/wiki/Yamaha_MU-series).

## Download and run
Download the latest binary from the [release page](https://github.com/ncirocco/midino/releases)

In the terminal run

`./midino -l` to list the available output MIDI devices and their IDs.

`./midino <device-id> <midi-file/playlist>` to play the selected midi file or a playlist.

If you want to be able to call **midino** from anywhere move the binary to the `/usr/bin/` directory.

`sudo cp midino /usr/bin/midino`

### Playlist
In order to make a playlist create a plain text file with one midi file per line using absolute paths to the files.

## Development requirements
Have GoLang installed and configured in your local environment. More information can be found [here](https://golang.org/doc/install).

### libportmidi-dev
`sudo apt install libportmidi-dev`
