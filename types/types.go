package types

type Trip struct {
	BooReturn      bool    `json:"boo_return"`
	CO2Kg          float64 `json:"co2_kg"`
	PriceEUR       float64 `json:"price_eur"`
	WorkTimeSec    float64 `json:"workTime_sec"`
	DurationOutSec float64 `json:"duration_out_sec"`
	Description    string  `json:"desc"`
	ID             string  `json:"id"`
	Score          float64 `json:"score,omitempty"`
}

type ByScore []Trip

func (s ByScore) Len() int {
	return len(s)
}

func (s ByScore) Swap(i,j int) {
	s[i],s[j] = s[j],s[i]
}

func (s ByScore) Less(i,j int) bool {
	return s[i].Score < s[j].Score
}
