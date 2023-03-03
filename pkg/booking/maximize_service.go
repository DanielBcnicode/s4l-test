package booking

import "sort"

type Maximizer interface {
	Maximize(bookings []Request) (MaximizeResponse, error)
}

type Maximize struct{}

func NewMaximizer() *Maximize {
	return &Maximize{}
}

func (m *Maximize) Maximize(bookings []Request) (MaximizeResponse, error) {
	sort.Slice(bookings, func(i, j int) bool {
		return bookings[i].StartDate().Unix() < bookings[j].StartDate().Unix()
	})
	return MaximizeResponse{}, nil
}
