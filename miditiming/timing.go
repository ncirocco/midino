package miditiming

// DefaultTempo is 500000 microseconds per beat for MIDI.
// Unlike music, tempo in MIDI is not given as beats per minute,
// but rather in microseconds per beat.
const DefaultTempo int64 = 500000

// CalculateBpm calculates and returns the BPM based on the given tempo
func CalculateBpm(tempo int64) int64 {
	return (60 * 1000000) / tempo
}

// TempoAndPpqnToNS calculates how many nanoseconds there are in one tick
// based on tempo and ppqn (pulses per quarter note)
func TempoAndPpqnToNS(tempo int64, ppqn int64) int64 {
	return tempo * 1000 / ppqn
}
