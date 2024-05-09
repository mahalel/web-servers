package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func validateHandler(w http.ResponseWriter, r *http.Request) {
	type reqBody struct {
		Body string `json:"body"`
	}

	type respBody struct {
		// the key will be the name of struct field unless you give it an explicit JSON tag
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(r.Body)
	params := reqBody{}

	err := decoder.Decode(&params)
	if err != nil {
		// an error will be thrown if the JSON is invalid or has the wrong types
		// any missing fields will simply have their values in the struct set to their zero value
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	} else {
		bodyLength := len(params.Body)
		fmt.Printf("%v", bodyLength)
	}

	// Response logic
	resp := respBody{}

	dat, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
	// params is a struct with data populated successfully
	// ...
}
