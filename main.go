package main

import (
	"log"

	"github.com/wlcsm/poll_app/handler"
	"github.com/wlcsm/poll_app/server"
)

func main() {

	s := server.New(8080)

	s.AddRoute("GET", "/poll/get", handler.QueryPoll)
	s.AddRoute("POST", "/poll/create", handler.CreatePoll)
	s.AddRoute("POST", "/poll/submit", handler.SubmitPoll)


	log.Println("Running on 127.0.0.1:8080")

	log.Fatal(s.Start())
}
