package booking

type StatsResponse struct {
	AverageNight float32 `json:"avg_night"`
	MinNight float32 `json:"min_night"`
	MaxNight float32 `json:"max_night"`
}
