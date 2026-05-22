package jmdict

import (
	"encoding/json"
	"io"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
)

type DictFile struct {
	Words []domain.Word `json:"words"`
}

func Parse(reader io.Reader) ([]*domain.Word, error) {
	var file DictFile
	if err := json.NewDecoder(reader).Decode(&file); err != nil {
		return nil, err
	}

	words := make([]*domain.Word, 0, len(file.Words))
	for i := range file.Words {
		w := &file.Words[i]
		w.Source = "admin"
		words = append(words, w)
	}

	return words, nil
}
