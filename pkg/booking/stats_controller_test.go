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

	req, err := http.NewRequest("POST", "/stats", strings.NewReader(`[
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
	]`))

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

	assert.Equal(t, float32(8.29), got.AverageNight)
	assert.Equal(t, float32(8), got.MinNight)
	assert.Equal(t, float32(8.58), got.MaxNight)

}
