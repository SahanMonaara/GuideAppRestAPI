package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Location Struct (Model)
type Location struct {
	ID  string `json:"id"`
	DID string `json:"did"`
	LAT string `json:"lat"`
	LNG string `json:"lng"`
}

var locations []Location

//Get all the locations
func getLocation(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)
}

//Create Location
func postLocation(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var location Location
	_ = json.NewDecoder(request.Body).Decode(&location)
	locations = append(locations, location)
	json.NewEncoder(w).Encode(location)
}

//Update the Location
func updateLocation(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range locations {
		if item.ID == params["id"] {
			locations = append(locations[:index], locations[index+1:]...)
			var location Location
			_ = json.NewDecoder(request.Body).Decode(&location)
			location.ID = params["id"]
			locations = append(locations, location)
			json.NewEncoder(w).Encode(location)
			return
		}
	}
	json.NewEncoder(w).Encode(locations)
}
func main() {
	//initial router

	router := mux.NewRouter()

	locations = append(locations, Location{ID: "1", DID: "765765765", LAT: "12.6788", LNG: "2.76576"})

	//Router Handlers
	router.HandleFunc("/api/locations", getLocation).Methods("GET")
	router.HandleFunc("/api/locations", postLocation).Methods("POST")
	router.HandleFunc("/api/locations/{id}", updateLocation).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))
}
