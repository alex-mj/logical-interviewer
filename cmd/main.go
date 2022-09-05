package main

import (
	"fmt"

	"github.com/alex-mj/logical-interview/service/api"
	interviewer "github.com/alex-mj/logical-interview/service/interviewer"
	"github.com/alex-mj/logical-interview/service/store"
)

func main() {
	fmt.Println("Hi, I'm logical interviewer >>> starting API (gRPC service)....")

	// TODO: add log
	// TODO: add config
	imStore := store.InMemory{}
	interviewer := interviewer.Interviewer{DB: imStore}
	api.NewAPIServer(interviewer)
}
