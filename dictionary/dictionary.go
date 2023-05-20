package dictionary

import (
	"bufio"
	"embed"
	"log"
	"math/rand"
	"strings"
	"time"
)

type Dictionary struct {
	history []string
	words   []string
	rng     *rand.Rand
}

//go:embed assets
var fs embed.FS

func New() *Dictionary {
	d := &Dictionary{}
	d.Load()
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	d.rng = rand.New(source)
	return d
}

func (d *Dictionary) Load() {
	f, err := fs.Open("assets/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var words []string
	for scanner.Scan() {
		words = append(words, strings.ToUpper(scanner.Text()))
	}
	d.words = words
}

func (d *Dictionary) Random() string {
	ii := d.rng.Intn(len(d.words))
	selected := d.words[ii]
	d.history = append(d.history, selected)
	return selected
}

func (d *Dictionary) ClearHistory() {
	d.history = []string{}
}

func (d *Dictionary) Length() int {
	return len(d.words)
}

func (d *Dictionary) History() []string {
	return d.history
}
