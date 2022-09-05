package api

import (
	"log"
	"net"

	domain "github.com/alex-mj/logical-interview/service"
	pb "github.com/alex-mj/logical-interview/service/api/proto/interview"
	"google.golang.org/grpc"
)

func NewAPIServer(Interviewer domain.Interviewer) InterviewServer {

	s := grpc.NewServer()
	InterviewServer := InterviewServer{Interviewer: Interviewer}
	pb.RegisterLogicalInterviewServer(s, InterviewServer)

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
