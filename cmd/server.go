package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/danielbcnicode/timeslot/pkg/booking"
	"github.com/gorilla/mux"
)

func main() {
	runMainApp()
}

func runMainApp() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Println("Goodby")
		os.Exit(1)
	}()

	payloadExtractor := booking.NewPayloadExtract()
	statsCalculator := booking.NewStatsCalculator()
	maximizer := booking.NewMaximizer()

	router := mux.NewRouter()
	router.HandleFunc("/stats", booking.StatsController(payloadExtractor, statsCalculator)).Methods("POST")
	router.HandleFunc("/maximize", booking.MaximizeController(payloadExtractor, maximizer)).Methods("POST")

	log.Println("Listen on port 8088")

	_ = http.ListenAndServe(":8088", router)

}
