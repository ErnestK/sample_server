package positions

import "time"

// Position need for deserialaizinf response
type Position struct {
	Keyword  string    `json:"keyword"`
	Position int       `json:"position"`
	URL      string    `json:"url"`
	Volume   int       `json:"volume"`
	Results  int       `json:"results"`
	Updated  time.Time `json:"updated"`
}

// DomainWithPosition need for deserialaizing response
type DomainWithPosition struct {
	Domain    string     `json:"domain"`
	Positions []Position `json:"positions"`
}
