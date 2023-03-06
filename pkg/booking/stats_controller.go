package booking

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// This file has the API stats controller
// It gets the payload and deserialize it in the proper struct
// The service to calculate the averages is injected in the constructor

var (
	TotalDaysCanBeZero = errors.New("total days can't be zero")
)

func StatsController(
	extractor PayloadExtractor,
	calculator StatsServiceCalculator,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bookings, err := extractor.ExtractPayload(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, err.Error())))
			return
		}

		retData, err := calculator.Calculate(bookings)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(retData)
	}
}
