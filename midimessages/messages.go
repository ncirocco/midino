package midimessages

// MidiMessage represents a message that can be sent to the MIDI device
type MidiMessage struct {
	StatusByte int64
	DataByte1  int64
	DataByte2  int64
}

// StopAll returns the midi message to stop all sounds for the given channel
func StopAll(channel int64) *MidiMessage {
	return &MidiMessage{
		StatusByte: 0xB0 + channel,
		DataByte1:  0x78,
		DataByte2:  0x00,
	}
}
