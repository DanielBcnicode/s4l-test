package booking

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatsController(t *testing.T) {
	service := StatsController()

	type args struct {
		payload string
	}
	tests := []struct {
		name string
		args args
		want StatsResponse
	} {
		{
			name: "two bookings",
			args: args{
				payload: `[
					{
						"request_id":"bookata_XY123",
						"check_in":"2020-01-01",
						"nights":5,
						"selling_rate":200,
						"margin":20
					},
					{
						"request_id":"kayete_PP234",
						"check_in":"2020-01-04",
						"nights":4,
						"selling_rate":156,
						"margin":22
					}
				]`,
			},
			want: StatsResponse{
				AverageNight: 8.29,
				MaxNight: 8.58,
				MinNight: 8,
			},
		},
		{
			name: "three bookings",
			args: args{
				payload: `[
					{
						"request_id":"bookata_XY123",
						"check_in":"2020-01-01",
						"nights":1,
						"selling_rate":50,
						"margin":20
					},
					{
						"request_id":"kayete_PP234",
						"check_in":"2020-01-04",
						"nights":1,
						"selling_rate":55,
						"margin":22
					},
					{
						"request_id":"trivoltio_ZX69",
						"check_in":"2020-01-07",
						"nights":1,
						"selling_rate":49,
						"margin":21
					}
				]`,
			},
			want: StatsResponse{
				AverageNight: 10.80,
				MaxNight: 12.1,
				MinNight: 10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/stats", strings.NewReader(tt.args.payload))
			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()
			handle := http.HandlerFunc(service)
			handle.ServeHTTP(recorder, req)

			res := recorder.Result()
			resBody, _ := io.ReadAll(res.Body)

			got := StatsResponse{}
			_ = json.Unmarshal(resBody, &got)

			assert.Equal(t, tt.want, got)
		})
	}

}
