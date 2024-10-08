package booking

// MaximizeResponse is the data structure to return in the Maximize end-point
type MaximizeResponse struct {
	RequestIDs   []string `json:"request_ids"`
	TotalProfit  float32  `json:"total_profit"`
	AverageNight float32  `json:"avg_night"`
	MinNight     float32  `json:"min_night"`
	MaxNight     float32  `json:"max_night"`
}
