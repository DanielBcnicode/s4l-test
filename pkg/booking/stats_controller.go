package booking

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"math"
	"net/http"
)

// This file has the API stats controller
// It gets the payload and deserialize it in the proper struct
// The service to calculate the averages is injected in the constructor

var (
	ErrorCantReadBody  = errors.New("can't read the request body")
	TotalDaysCanBeZero = errors.New("total days can't be zero")
)

func StatsController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		d, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: write the correct answer in the body
			return
		}

		data := []RequestAPI{}
		err = json.Unmarshal(d, &data)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: write the correct answer in the body
			return
		}

		bookings := make([]Request, 0)
		for _, r := range data {
			reqAPI, err := RequestFromRequestAPI(r)
			if err != nil {
				log.Printf("ERROR: %s\n", err)
				w.WriteHeader(http.StatusBadRequest)
				// TODO: write the correct answer in the body
				return
			}

			bookings = append(bookings, reqAPI)
		}

		retData, err := StatsService(bookings)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(retData)
	}
}

// StatsService returns the profit calculus for the bookings passed as parameter
func StatsService(bookings []Request) (StatsResponse, error) { //Put this service in one object with interface and inject it
	min, max, total := float32(0), float32(0), float32(0)
	totalBookings := 0

	for i, r := range bookings {
		if i == 0 {
			min = r.Profit
			max = r.Profit
		}

		if r.Profit > max {
			max = r.Profit
		}
		if r.Profit < min {
			min = r.Profit
		}
		total += r.Profit
		totalBookings += 1

	}

	if totalBookings == 0 {
		return StatsResponse{}, TotalDaysCanBeZero
	}

	return StatsResponse{
		AverageNight: floatRoundPrecision(total / float32(totalBookings), 2),
		MinNight:     min,
		MaxNight:     max,
	}, nil
}

func floatRoundPrecision(n float32, p int) float32 {
	return float32(math.Round(float64(n)*(float64(math.Pow10(p)))) / math.Pow10(p))
}
