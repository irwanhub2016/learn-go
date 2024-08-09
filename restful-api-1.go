package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Profile struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Hobby string `json:"hobby"`
}

type Profiles []Profile

func allProfiles(w http.ResponseWriter, r *http.Request) {
	profiles := Profiles{
		Profile{Name: "Irwan", Age: 25, Hobby: "Coding"},
	}
	fmt.Println("Endpoint Hit: All Profiles Endpoint")
	json.NewEncoder(w).Encode(profiles)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		invalidEndpoint(w, r)
		return
	}
	fmt.Fprintf(w, "Hello World Irwan")
}

func invalidEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "your endpoint invalid")
}

func handleRequest() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/api/v1/profiles", allProfiles)

	fmt.Println("Server running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}