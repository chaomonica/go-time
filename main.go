//https://tutorialedge.net/golang/creating-restful-api-with-golang/
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//handle all requests to root URL
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homePage")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	utc, err := time.LoadLocation("America/New_York")
	// utc := time.Now().UTC().Format("Jan 02, 2006 03:04 PM")
	// cst := time.FixedZone("America/Chicago", -6*3600)

	if err != nil {
		fmt.Fprintln(w, "err: ", err.Error())
	}
	fmt.Fprintln(w, "Location:", utc, ":Time:", t.In(utc))

	// fmt.Fprint(w, "Current time is...", utc, "and", cst)
	// fmt.Println("Current time!")
}

func timeSingleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["zone"]

	fmt.Fprintf(w, "Key: "+key)
}

//match URL path hit with a defined function
func handleRequests() {
	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("time/{zone}", timeSingleHandler)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

//kick off the API
func main() {
	handleRequests()
}

//comment to test ssh
