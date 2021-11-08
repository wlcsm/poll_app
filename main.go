package main

import (
	"log"
	"net/http"

	"github.com/poll_app/handler"
)

func main() {

	http.HandleFunc("/poll/", handler.GetPoll)
	http.HandleFunc("/create/", handler.CreatePoll)
	http.HandleFunc("/submit/", handler.SubmitPoll)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
