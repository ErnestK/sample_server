package summary

// CountByDomain need for deserialaizing response
type CountByDomain struct {
	Domain string `json:"domain"`
	Count  int    `json:"positions_count"`
}
