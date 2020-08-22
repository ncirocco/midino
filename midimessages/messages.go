package midimessages

const (
	// AcousticGrandPiano instrument
	AcousticGrandPiano = iota
	// BrightAcousticPiano instrument
	BrightAcousticPiano
	// ElectricGrandPiano instrument
	ElectricGrandPiano
	// HonkyTonkPiano instrument
	HonkyTonkPiano
	// ElectricPiano1 instrument
	ElectricPiano1
	// ElectricPiano2 instrument
	ElectricPiano2
	// Harpsichord instrument
	Harpsichord
	// Clavi instrument
	Clavi
	// Celesta instrument
	Celesta
	// Glockenspiel instrument
	Glockenspiel
	// MusicBox instrument
	MusicBox
	// Vibraphone instrument
	Vibraphone
	// Marimba instrument
	Marimba
	// Xylophone instrument
	Xylophone
	// TubularBells instrument
	TubularBells
	// Dulcimer instrument
	Dulcimer
	// DrawbarOrgan instrument
	DrawbarOrgan
	// PercussiveOrgan instrument
	PercussiveOrgan
	// RockOrgan instrument
	RockOrgan
	// ChurchOrgan instrument
	ChurchOrgan
	// ReedOrgan instrument
	ReedOrgan
	// Accordion instrument
	Accordion
	// Harmonica instrument
	Harmonica
	// TangoAccordion instrument
	TangoAccordion
	// AcousticGuitarNylon instrument
	AcousticGuitarNylon
	// AcousticGuitarSteel instrument
	AcousticGuitarSteel
	// ElectricGuitarJazz instrument
	ElectricGuitarJazz
	// ElectricGuitarClean instrument
	ElectricGuitarClean
	// ElectricGuitarMuted instrument
	ElectricGuitarMuted
	// OverdrivenGuitar instrument
	OverdrivenGuitar
	// DistortionGuitar instrument
	DistortionGuitar
	// GuitarHarmonics instrument
	GuitarHarmonics
	// AcousticBass instrument
	AcousticBass
	// ElectricBassFinger instrument
	ElectricBassFinger
	// ElectricBassPick instrument
	ElectricBassPick
	// FretlessBass instrument
	FretlessBass
	// SlapBass1 instrument
	SlapBass1
	// SlapBass2 instrument
	SlapBass2
	// SynthBass1 instrument
	SynthBass1
	// SynthBass2 instrument
	SynthBass2
	// Violin instrument
	Violin
	// Viola instrument
	Viola
	// Cello instrument
	Cello
	// Contrabass instrument
	Contrabass
	// TremoloStrings instrument
	TremoloStrings
	// PizzicatoStrings instrument
	PizzicatoStrings
	// OrchestralHarp instrument
	OrchestralHarp
	// Timpani instrument
	Timpani
	// StringEnsemble1 instrument
	StringEnsemble1
	// StringEnsemble2 instrument
	StringEnsemble2
	// SynthStrings1 instrument
	SynthStrings1
	// SynthStrings2 instrument
	SynthStrings2
	// ChoirAahs instrument
	ChoirAahs
	// VoiceOohs instrument
	VoiceOohs
	// SynthVoice instrument
	SynthVoice
	// OrchestraHit instrument
	OrchestraHit
	// Trumpet instrument
	Trumpet
	// Trombone instrument
	Trombone
	// Tuba instrument
	Tuba
	// MutedTrumpet instrument
	MutedTrumpet
	// FrenchHorn instrument
	FrenchHorn
	// BrassSection instrument
	BrassSection
	// SynthBrass1 instrument
	SynthBrass1
	// SynthBrass2 instrument
	SynthBrass2
	// SopranoSax instrument
	SopranoSax
	// AltoSax instrument
	AltoSax
	// TenorSax instrument
	TenorSax
	// BaritoneSax instrument
	BaritoneSax
	// Oboe instrument
	Oboe
	// EnglishHorn instrument
	EnglishHorn
	// Bassoon instrument
	Bassoon
	// Clarinet instrument
	Clarinet
	// Piccolo instrument
	Piccolo
	// Flute instrument
	Flute
	// Recorder instrument
	Recorder
	// PanFlute instrument
	PanFlute
	// Blownbottle instrument
	Blownbottle
	// Shakuhachi instrument
	Shakuhachi
	// Whistle instrument
	Whistle
	// Ocarina instrument
	Ocarina
	// Lead1Square instrument
	Lead1Square
	// Lead2Sawtooth instrument
	Lead2Sawtooth
	// Lead3Calliope instrument
	Lead3Calliope
	// Lead4Chiff instrument
	Lead4Chiff
	// Lead5Charang instrument
	Lead5Charang
	// Lead6Voice instrument
	Lead6Voice
	// Lead7Fifths instrument
	Lead7Fifths
	// Lead8BassLead instrument
	Lead8BassLead
	// Pad1Newage instrument
	Pad1Newage
	// Pad2Warm instrument
	Pad2Warm
	// Pad3Polysynth instrument
	Pad3Polysynth
	// Pad4Choir instrument
	Pad4Choir
	// Pad5Bowed instrument
	Pad5Bowed
	// Pad6Metallic instrument
	Pad6Metallic
	// Pad7Halo instrument
	Pad7Halo
	// Pad8Sweep instrument
	Pad8Sweep
	// FX1Rain instrument
	FX1Rain
	// FX2Soundtrack instrument
	FX2Soundtrack
	// FX3Crystal instrument
	FX3Crystal
	// FX4Atmosphere instrument
	FX4Atmosphere
	// FX5Brightness instrument
	FX5Brightness
	// FX6Goblins instrument
	FX6Goblins
	// FX7Echoes instrument
	FX7Echoes
	// FX8SciFi instrument
	FX8SciFi
	// Sitar instrument
	Sitar
	// Banjo instrument
	Banjo
	// Shamisen instrument
	Shamisen
	// Koto instrument
	Koto
	// Kalimba instrument
	Kalimba
	// Bagpipe instrument
	Bagpipe
	// Fiddle instrument
	Fiddle
	// Shanai instrument
	Shanai
	// TinkleBell instrument
	TinkleBell
	// Agogo instrument
	Agogo
	// SteelDrums instrument
	SteelDrums
	// Woodblock instrument
	Woodblock
	// TaikoDrum instrument
	TaikoDrum
	// MelodicTom instrument
	MelodicTom
	// SynthDrum instrument
	SynthDrum
	// ReverseCymbal instrument
	ReverseCymbal
	// GuitarFretNoise instrument
	GuitarFretNoise
	// BreathNoise instrument
	BreathNoise
	// Seashore instrument
	Seashore
	// BirdTweet instrument
	BirdTweet
	// TelephoneRing instrument
	TelephoneRing
	// Helicopter instrument
	Helicopter
	// Applause instrument
	Applause
	// Gunshot instrument
	Gunshot
)

const channelMode = 0xB0
const controlChange = 0xC0

// MidiMessage represents a message that can be sent to the MIDI device
type MidiMessage struct {
	StatusByte int64
	DataByte1  int64
	DataByte2  int64
}

// StopAll returns the midi message to stop all sounds for the given channel immidiately
func StopAll(channel int64) *MidiMessage {
	return &MidiMessage{
		StatusByte: channelMode + channel,
		DataByte1:  0x78,
		DataByte2:  0x00,
	}
}

// MuteAll mutes all the notes that are currently playing
func MuteAll(channel int64) *MidiMessage {
	return &MidiMessage{
		StatusByte: channelMode + channel,
		DataByte1:  0x7B,
		DataByte2:  0x00,
	}
}

// SelectInstrument returns the midi message to select the given instrument for the given channel
func SelectInstrument(channel int64, instrument int64) *MidiMessage {
	return &MidiMessage{
		StatusByte: controlChange + channel,
		DataByte1:  0x00,
		DataByte2:  0x00,
	}
}

// ClearAllControllersForChannel clears all the controllers to their default values for the given channel
func ClearAllControllersForChannel(channel int64) *MidiMessage {
	return &MidiMessage{
		StatusByte: controlChange + channel,
		DataByte1:  0x79,
		DataByte2:  0x00,
	}
}
