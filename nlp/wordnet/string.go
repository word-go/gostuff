package wordnet

import (
	"bytes"
	"fmt"
)

// String functions for types.

// TODO(amit): Consider removing String() functions to simplify the API.

func (w *Wordnet) String() string {
	return fmt.Sprintf("Wordnet[%d lemmas, %d synsets, %d exceptions]",
		len(w.Lemma), len(w.Synset), len(w.Exception))
}

func (s *Synset) String() string {
	result := bytes.NewBuffer(make([]byte, 0, 100))
	fmt.Fprintf(result, "Synset[%s.", s.Pos)
	for i, word := range s.Word {
		if i > 0 {
			fmt.Fprintf(result, ",")
		}
		fmt.Fprintf(result, " %v", word)
	}
	fmt.Fprintf(result, ": %s]", s.Gloss)
	return result.String()
}
