package booking

import (
	"time"

	"github.com/danielbcnicode/timeslot/internal"
)

type RequestAPI struct {
	RequestID   string `json:"request_id"`
	CheckIn     string `json:"check_in"`
	Nights      uint32 `json:"nights"`
	SellingRate uint32 `json:"selling_rate"`
	Margin      uint32 `json:"margin"`
}

type Request struct {
	ID          string
	SellingRate uint32
	Margin      uint32
	Profit      float32
	internal.DaySlot
}

func RequestFromRequestAPI(req RequestAPI) (Request, error) {
	t, err := time.Parse("2006-01-02", req.CheckIn)
	if err != nil {
		return Request{}, err
	}
	d := internal.NewDaySlot(t, t.Add(24*time.Hour*time.Duration(req.Nights)))
	return Request{
		req.RequestID,
		req.SellingRate,
		req.Margin,
		float32(req.SellingRate) * float32(req.Margin/100) / float32(req.Nights),
		d,
	}, nil
}
