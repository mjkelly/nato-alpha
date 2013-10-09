// alphabets.go handles loading phonetic alphabets and translating strings.
package translate

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

// PhoneticAlphabet represents a particular phonetic alphabet, with
// translations for each character.
type PhoneticAlphabet struct {
	Type    string
	Letters map[rune]Translation
}

// Translation represents an individual element of phonetic alphabet, with the
// name (e.g., "Charlie") and a pronundication (e.g., "CHAR-lee").
type Translation struct {
	Name, Pronunciation string
}

// MustLoad is wrapper around Load which panics on error.
func MustLoad(file string, name string) *PhoneticAlphabet {
	alpha, err := Load(file, name)
	if err != nil {
		panic(err)
	}
	return alpha
}

// Load loads the named CSV file representing a phonetic
// alphabet and returns a *PhoneticAlphabet.
func Load(file string, name string) (*PhoneticAlphabet, error) {
	alpha := &PhoneticAlphabet{Type: name}
	alpha.Letters = make(map[rune]Translation)

	fh, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	reader := csv.NewReader(fh)
	reader.Comment = '#'
	reader.LazyQuotes = true

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if len(line) != 3 {
			return nil, fmt.Errorf("csv line %v has %d parts, expected %d", line, len(line), 3)
		}
		if len(line[0]) != 1 {
			return nil, fmt.Errorf("field %d must be a single character, found %d", 0, len(line[0]))
		}
		key := unicode.ToLower(rune(line[0][0]))
		alpha.Letters[key] = Translation{Name: line[1], Pronunciation: line[2]}
	}
	log.Printf("Loaded %d translations", len(alpha.Letters))
	return alpha, nil
}

// Get returns the translation for a single rune, or nil if there is none in
// the receiving PhoneticAlphabet.
//
// The rune should be the lowercase version of its letter, if applicable.
func (pa *PhoneticAlphabet) Get(letter rune) *Translation {
	t, found := pa.Letters[unicode.ToLower(letter)]
	if found {
		return &t
	} else {
		return nil
	}
}

// Translate translates an entire phrase and returns the corresponding Translation objects.
//
// Any rune in the phrase that doesn't have a translation is replaced with a
// substitute Translation indicating that.
func (pa *PhoneticAlphabet) Translate(phrase string) []*Translation {
	var t []*Translation
	for _, letter := range phrase {
		if !unicode.IsSpace(letter) {
			if trans := pa.Get(letter); trans != nil {
				t = append(t, trans)
			} else {
				t = append(t, &Translation{Name: fmt.Sprintf("\"%s\"", string(letter)), Pronunciation: "No translation"})
			}
		}
	}
	return t
}
