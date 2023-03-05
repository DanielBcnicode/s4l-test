package booking

import (
	"errors"
	"time"

	"github.com/danielbcnicode/timeslot/internal"
)

var (
	ErrorNightsCanBeZero      = errors.New("nights can not be zero")
	ErrorSellingRateCanBeZero = errors.New("sellingRate can not be zero")
)

type RequestAPI struct {
	RequestID   string `json:"request_id"`
	CheckIn     string `json:"check_in"`
	Nights      uint32 `json:"nights"`
	SellingRate uint32 `json:"selling_rate"`
	Margin      uint32 `json:"margin"`
}

type Request struct {
	ID             string
	SellingRate    uint32
	Margin         uint32
	ProfitPerNight float32
	Profit         float32
	internal.DaySlot
}

func RequestFromRequestAPI(req RequestAPI) (Request, error) {
	t, err := time.Parse("2006-01-02", req.CheckIn)
	if err != nil {
		return Request{}, err
	}

	if req.Nights == 0 {
		return Request{}, ErrorNightsCanBeZero
	}

	if req.SellingRate == 0 {
		return Request{}, ErrorSellingRateCanBeZero
	}

	d := internal.NewDaySlot(t, t.Add(24*time.Hour*time.Duration(req.Nights)))
	return Request{
		req.RequestID,
		req.SellingRate,
		req.Margin,
		float32(req.SellingRate) * (float32(req.Margin) / float32(100)) / float32(req.Nights),
		float32(req.SellingRate) * (float32(req.Margin) / float32(100)),
		d,
	}, nil
}
