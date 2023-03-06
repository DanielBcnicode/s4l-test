package booking

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// This file has the API maximize controller
// It gets the payload and deserialize it in the proper struct

func MaximizeController(
	extractor PayloadExtractor,
	maximizer Maximizer,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bookings, err := extractor.ExtractPayload(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, err.Error())))
			return
		}

		retData := maximizer.Maximize(bookings)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(retData)
	}
}
