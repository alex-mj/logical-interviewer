package main

import (
	"fmt"

	"github.com/alex-mj/logical-interview/service/api"
	interviewer "github.com/alex-mj/logical-interview/service/interviewer"
	"github.com/alex-mj/logical-interview/service/store"
)

func main() {
	fmt.Println("Hi, I'm logical gRPC service")

	imStore := store.InMemory{}
	Interviewer := interviewer.Interviewer{DB: imStore}
	api.NewAPIServer(Interviewer)
}
