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
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Println("Goodbye")
		os.Exit(1)
	}()

	payloadExtractor := booking.NewPayloadExtract()
	statsCalculator := booking.NewStatsCalculator()
	maximizer := booking.NewMaximizer()

	router := mux.NewRouter()
	router.HandleFunc("/stats", booking.StatsController(payloadExtractor, statsCalculator)).Methods("POST")
	router.HandleFunc("/maximize", booking.MaximizeController(payloadExtractor, maximizer)).Methods("POST")

	serverAddress := configServer()
	log.Printf("Listen on %s \n", serverAddress)

	err := http.ListenAndServe(serverAddress, router)
	if err != nil {
		log.Fatal("Error in the HTTP server :" + err.Error())
	}
}

// configServer Get the Server configuration from the Environment variable SRV_PORT
func configServer() string {
	defaultAddress := ":8088"

	configAddress := ":" + os.Getenv("SRV_PORT")
	if configAddress == ":" {
		configAddress = defaultAddress
	}

	return configAddress
}
