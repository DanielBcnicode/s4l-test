package booking

import (
	"errors"
	"time"

	"github.com/danielbcnicode/timeslot/internal"
)

var (
	ErrorNightsCantBeZero      = errors.New("ERROR: nights can not be zero or empty")
	ErrorSellingRateCantBeZero = errors.New("ERROR: sellingRate can not be zero or empty")
	ErrorDateFormatWrong       = errors.New("ERROR: check_in format is wrong")
	ErrorIDCantBeEmpty         = errors.New("ERROR: request_id can not be empty")
	ErrorMarginCantBeZero      = errors.New("ERROR: margin can not be zero or empty")
)

// RequestAPI is the data structure used as parameters in the 2 end-points
type RequestAPI struct {
	RequestID   string `json:"request_id"`
	CheckIn     string `json:"check_in"`
	Nights      uint32 `json:"nights"`
	SellingRate uint32 `json:"selling_rate"`
	Margin      uint32 `json:"margin"`
}

// Request is the data transformed from the request pay-load to be used in the system
type Request struct {
	ID             string
	SellingRate    uint32
	Margin         uint32
	ProfitPerNight float32
	Profit         float32
	internal.DaySlot
}

// RequestFromRequestAPI is the domain Request constructor from the API request
func RequestFromRequestAPI(req RequestAPI) (Request, error) {
	t, err := time.Parse("2006-01-02", req.CheckIn)
	if err != nil {
		return Request{}, ErrorDateFormatWrong
	}
	if req.RequestID == "" {
		return Request{}, ErrorIDCantBeEmpty
	}
	if req.Margin == 0 {
		return Request{}, ErrorMarginCantBeZero
	}
	if req.Nights == 0 {
		return Request{}, ErrorNightsCantBeZero
	}

	if req.SellingRate == 0 {
		return Request{}, ErrorSellingRateCantBeZero
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
