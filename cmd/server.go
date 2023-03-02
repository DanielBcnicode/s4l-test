package main

import (
	"net/http"

	"github.com/danielbcnicode/timeslot/pkg/booking"
	"github.com/gorilla/mux"
)

func main() {
	runMainApp()
}

func runMainApp() {
	router := mux.NewRouter()
	router.HandleFunc("/stats", booking.StatsController()).Methods("POST")
	_ = http.ListenAndServe(":8088", router)

}
