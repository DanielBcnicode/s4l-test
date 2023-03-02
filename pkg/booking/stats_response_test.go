package booking

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatsResponseSerialization(t *testing.T) {
	serialized := []byte(`{
		"avg_night":8.29,
		"min_night":8,
		"max_night":8.58
		}`)
	want := StatsResponse{
		AverageNight: 8.29,
		MinNight:     8,
		MaxNight:     8.58,
	}

	got := StatsResponse{}
	if err := json.Unmarshal(serialized, &got); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want, got)

	serializedGot, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	assert.JSONEq(t, string(serialized), string(serializedGot))

}
