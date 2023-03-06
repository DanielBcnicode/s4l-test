package booking

import (
	"encoding/json"
	"io"
	"net/http"
)

// PayloadExtractor is the interface to the Payload extractor service
type PayloadExtractor interface {
	ExtractPayload(r *http.Request) ([]Request, error)
}

// ExtractPayload extracts the payload for the 2 endpoints

// PayloadExtract the implementation of the service
type PayloadExtract struct {
}

// NewPayloadExtract creates a new service
func NewPayloadExtract() *PayloadExtract {
	return &PayloadExtract{}
}

// ExtractPayload has the service main logic. It extracts the payload and returns
// the data in the proper format to be used in the system
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
