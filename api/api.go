package api

import (
	"context"
	"fmt"

	domain "github.com/alex-mj/logical-interview/service"
	pb "github.com/alex-mj/logical-interview/service/api/proto/interview"
)

// InterviewServer is a server that exposes the logical Interview service.
type InterviewServer struct {
	pb.UnimplementedLogicalInterviewServer
	Interviewer domain.Interviewer
}

// GetScores - get last test result
func (InterviewServer) GetResult(ctx context.Context, user *pb.User) (*pb.Result, error) {
	fmt.Println("user.Id:", user.Id, "TODO: implement InterviewServer.GetResult()")
	return &pb.Result{Scores: 30, NewAttempInXDay: 29}, nil
}

// GetScores - get new question
func (InterviewServer) GetQuestion(ctx context.Context, user *pb.User) (*pb.Question, error) {
	fmt.Println("user.Id:", user.Id, "TODO: implement InterviewServer.GetQuestion()")
	return &pb.Question{
		Id:        "WE3WR25WERNYW",
		Condition: "Если тряхнуть бурдылькой, то начнется стрельба. Бурдылькой тряхнули.",
		QuestionOption: []*pb.QuestionOption{
			{Id: "1Q1", Data: "стрельба уже началась;"},
			{Id: "1Q2", Data: "стрельба начнется когда-нибудь;"},
			{Id: "1Q3", Data: "стрельба начнется когда-нибудь или уже началась;"},
		},
	}, nil
}

// service LogicalInterview {
//		TODO:
//     rpc SendAnswer(Ansver) returns (AnswerResponce);
//     rpc GetSummary(User) returns (Summary);
// }
