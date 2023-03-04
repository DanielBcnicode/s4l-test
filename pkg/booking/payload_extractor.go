package booking

import (
	"encoding/json"
	"io"
	"net/http"
)

type PayloadExtractor interface {
	ExtractPayload(r *http.Request) ([]Request, error)
}

// ExtractPayload extracts the payload for the 2 endpoints

type PayloadExtract struct {
}

func NewPayloadExtract() *PayloadExtract {
	return &PayloadExtract{}
}

func (p *PayloadExtract) ExtractPayload(r *http.Request) ([]Request, error) {
	d, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var data []RequestAPI
	err = json.Unmarshal(d, &data)
	if err != nil {
		return nil, err
	}

	bookings := make([]Request, 0)
	for _, r := range data {
		reqAPI, err := RequestFromRequestAPI(r)
		if err != nil {
			return nil, err
		}

		bookings = append(bookings, reqAPI)
	}

	return bookings, nil
}
