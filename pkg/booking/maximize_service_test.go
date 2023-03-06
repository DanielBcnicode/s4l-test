package booking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximize_Maximize(t *testing.T) {
	r1, _ := RequestFromRequestAPI(RequestAPI{
		RequestID:   "bookata_XY123",
		CheckIn:     "2020-01-01",
		Nights:      5,
		SellingRate: 200,
		Margin:      20,
	})
	r2, _ := RequestFromRequestAPI(RequestAPI{
		RequestID:   "kayete_PP234",
		CheckIn:     "2020-01-04",
		Nights:      4,
		SellingRate: 156,
		Margin:      5,
	})
	r3, _ := RequestFromRequestAPI(RequestAPI{
		RequestID:   "atropote_AA930",
		CheckIn:     "2020-01-04",
		Nights:      4,
		SellingRate: 150,
		Margin:      6,
	})
	r4, _ := RequestFromRequestAPI(RequestAPI{
		RequestID:   "acme_AAAAA",
		CheckIn:     "2020-01-10",
		Nights:      4,
		SellingRate: 160,
		Margin:      30,
	})

	type args struct {
		bookings []Request
	}
	tests := []struct {
		name string
		args args
		want MaximizeResponse
	}{
		{
			name: "Happy path",
			args: args{
				bookings: []Request{r4, r3, r2, r1}, // Change the order to test the sort function
			},
			want: MaximizeResponse{
				RequestIDs:   []string{"bookata_XY123", "acme_AAAAA"},
				TotalProfit:  88,
				AverageNight: 10,
				MinNight:     8,
				MaxNight:     12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Maximize{}
			got := m.Maximize(tt.args.bookings)
			assert.Equalf(t, tt.want, got, "Maximize(%v)", tt.args.bookings)
		})
	}
}
