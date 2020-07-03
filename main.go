package main

import (
	"fmt"
	"net/http"
	"math/rand"
)

var randomNumber int

func main() {
	getRand()
	http.HandleFunc("/rand", handleRand)
	http.HandleFunc("/", sayRand)
	http.ListenAndServe(":8080", nil)
}

func getRand() {
	randomNumber = rand.Int()
}

func handleRand(w http.ResponseWriter, r *http.Request) {
	getRand()
}

func sayRand(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, randomNumber)
}
