package sound

import (
	"embed"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const (
	SampleRate = 48000
	SoundsDir  = "assets/sounds"
)

var soundMgr *SoundManager

//go:embed assets
var fs embed.FS

type SoundManager struct {
	AudioContext *audio.Context
	SFXLibrary   map[string]*audio.Player
}

func New() *SoundManager {
	mgr := &SoundManager{}
	mgr.SFXLibrary = make(map[string]*audio.Player)
	return mgr
}

func GetSoundManager() *SoundManager {
	if soundMgr != nil {
		return soundMgr
	}

	return New()
}

func (s *SoundManager) Load() {
	s.AudioContext = audio.CurrentContext()

	if s.AudioContext == nil {
		s.AudioContext = audio.NewContext(SampleRate)
	}

	s.AddSound("text-input", "text_1_16.wav")
	s.AddSound("text-delete", "cancel_1_16.wav")
	s.AddSound("error", "boss_hit_1_16.wav")
	s.AddSound("good", "confirm_1_16.wav")
}

func (s *SoundManager) AddSound(key, fn string) {
	snd := MustLoadSound(fn)
	p, err := s.AudioContext.NewPlayer(snd)
	if err != nil {
		log.Panicf(err.Error())
	}
	s.SFXLibrary[key] = p
}

func (s *SoundManager) Play(name string) {
	s.SFXLibrary[name].Rewind()
	s.SFXLibrary[name].Play()
}

func MustLoadSound(n string) *wav.Stream {
	tf := fmt.Sprintf("%s/%s", SoundsDir, n)
	f, err := fs.Open(tf)
	if err != nil {
		log.Panicf("Could not open file %s", tf)
	}
	d, err := wav.DecodeWithoutResampling(f)
	if err != nil {
		log.Panicf("Could not decode file %s: %s", tf, err.Error())
	}
	return d
}
