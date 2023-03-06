package booking

// StatsResponse is the struct used to the stats end-point response
type StatsResponse struct {
	AverageNight float32 `json:"avg_night"`
	MinNight     float32 `json:"min_night"`
	MaxNight     float32 `json:"max_night"`
}
