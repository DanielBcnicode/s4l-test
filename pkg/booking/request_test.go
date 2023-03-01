package booking

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/danielbcnicode/timeslot/internal"
	"github.com/stretchr/testify/assert"
)

func TestRequstSerialization(t *testing.T) {
	serialized := `[
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
	]`

	want := []RequestAPI{
		{
			RequestID:   "bookata_XY123",
			CheckIn:     "2020-01-01",
			Nights:      5,
			SellingRate: 200,
			Margin:      20,
		},
		{
			RequestID:   "kayete_PP234",
			CheckIn:     "2020-01-04",
			Nights:      4,
			SellingRate: 156,
			Margin:      22,
		},
	}

	request := []RequestAPI{}

	err := json.Unmarshal([]byte(serialized), &request)
	if err != nil {
		t.Errorf("error unmarshalling json: %s\n", err)
	}

	assert.Equal(t, want, request)
	fmt.Printf("request := %#v \n", request)

}

func TestRequst(t *testing.T) {
	b := internal.NewDaySlot(time.Now(), time.Now().Add(24 * time.Hour))
	a := Request{}

	assert.Equal(t, a.Duration(), 1)
}