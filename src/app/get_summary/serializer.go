package get_summary

type CountByDomain struct {
    Domain  string `json:"domain"`
    Count   int    `json:"positions_count"`
}