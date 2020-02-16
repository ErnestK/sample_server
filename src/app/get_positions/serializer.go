package get_positions

import "time"

type Position struct {
	Keyword  string    `json:"keyword"`
	Position int       `json:"position"`
	Url      string    `json:"url"`
	Volume   int       `json:"volume"`
	Results  int       `json:"results"`
	Updated  time.Time `json:"updated"`
}

type DomainWithPosition struct {
	Domain    string     `json:"domain"`
	Positions []Position `json:"positions"`
}
