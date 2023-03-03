package booking

import "github.com/danielbcnicode/timeslot/internal"

type StatsServiceCalculator interface {
	Calculate(bookings []Request) (StatsResponse, error)
}

type StatsCalculator struct{}

func NewStatsCalculator() *StatsCalculator {
	return &StatsCalculator{}
}

// Calculate returns the profit calculus for the bookings passed as parameter
func (s *StatsCalculator) Calculate(bookings []Request) (StatsResponse, error) { //Put this service in one object with interface and inject it
	min, max, total := float32(0), float32(0), float32(0)
	totalBookings := 0

	for i, r := range bookings {
		if i == 0 {
			min = r.Profit
			max = r.Profit
		}

		if r.Profit > max {
			max = r.Profit
		}
		if r.Profit < min {
			min = r.Profit
		}
		total += r.Profit
		totalBookings += 1

	}

	if totalBookings == 0 {
		return StatsResponse{}, TotalDaysCanBeZero
	}

	return StatsResponse{
		AverageNight: internal.FloatRoundPrecision(total/float32(totalBookings), 2),
		MinNight:     min,
		MaxNight:     max,
	}, nil
}
