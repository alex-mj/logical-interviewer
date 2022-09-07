package api

import (
	"log"
	"net"

	interview "github.com/alex-mj/logical-interview/service"
	pb "github.com/alex-mj/logical-interview/service/api/proto/interview"
	"google.golang.org/grpc"
)

func NewAPIServer(intwr interview.Interviewer) ApiServer {

	s := grpc.NewServer()
	InterviewServer := ApiServer{Interviewer: intwr}
	pb.RegisterLogicalInterviewServer(s, &InterviewServer)

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
	return InterviewServer
}
