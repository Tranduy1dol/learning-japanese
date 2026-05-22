package domain

import "strings"

type Word struct {
	ID        string    `bson:"_id" json:"id"`
	EntSeq    string    `bson:"ent_seq" json:"ent_seq,omitempty"`
	Kanji     []Kanji   `bson:"kanji" json:"kanji"`
	Readings  []Reading `bson:"readings" json:"readings"`
	Senses    []Sense   `bson:"senses" json:"senses"`
	JLPT      int       `bson:"jlpt" json:"jlpt,omitempty"`
	IsCommon  bool      `bson:"is_common" json:"is_common,omitempty"`
	Source    string    `bson:"source" json:"source,omitempty"`
	CreatedBy string    `bson:"created_by" json:"created_by,omitempty"`
}

type Kanji struct {
	Text     string `bson:"text" json:"text"`
	Info     string `bson:"info" json:"info,omitempty"`
	Priority int    `bson:"priority" json:"priority,omitempty"`
}

type Reading struct {
	Text     string   `bson:"text" json:"text"`
	Status   string   `bson:"status" json:"status,omitempty"`
	Info     []string `bson:"info" json:"info,omitempty"`
	Priority int      `bson:"priority" json:"priority,omitempty"`
}

type Sense struct {
	POS    []string `bson:"pos" json:"pos"`
	Gloss  []Gloss  `bson:"gloss" json:"gloss"`
	Source string   `bson:"source" json:"source,omitempty"`
}

type Gloss struct {
	Text string `bson:"text" json:"text"`
	Lang string `bson:"lang" json:"lang"`
}

func (w *Word) SearchText() (title string, text string) {
	if len(w.Kanji) > 0 {
		title = w.Kanji[0].Text
	} else if len(w.Readings) > 0 {
		title = w.Readings[0].Text
	}

	var parts []string
	for _, k := range w.Kanji {
		parts = append(parts, k.Text)
	}
	for _, r := range w.Readings {
		parts = append(parts, r.Text)
	}
	for _, s := range w.Senses {
		for _, g := range s.Gloss {
			parts = append(parts, g.Text)
		}
	}

	text = strings.Join(parts, " ")
	return
}
