package api

import (
	"context"
	"fmt"

	domain "github.com/alex-mj/logical-interview/service"
	interview "github.com/alex-mj/logical-interview/service"
	pb "github.com/alex-mj/logical-interview/service/api/proto/interview"
)

// ApiServer is a server that exposes the logical Interview service.
type ApiServer struct {
	pb.UnimplementedLogicalInterviewServer
	Interviewer domain.Interviewer
}

// GetResult - get last test result
func (as *ApiServer) GetResult(ctx context.Context, user *pb.User) (*pb.Result, error) {
	fmt.Println("user.Id:", user.Id, "TODO: implement InterviewServer.GetResult()")
	return &pb.Result{Scores: 30, NewAttempInXDay: 29}, nil
}

// GetNextQuestion - get new question
//
func (as *ApiServer) GetNextQuestion(ctx context.Context, user *pb.User) (*pb.Question, error) {
	fmt.Println("user.Id:", user.Id, "TODO: implement InterviewServer.GetQuestion()")
	q := as.Interviewer.GetNextQuestion(interview.UserID(user.Id))

	pbOptions := make([]*pb.QuestionOption, len(q.QuestionOptions))
	for _, option := range q.QuestionOptions {
		pbOptions = append(pbOptions, &pb.QuestionOption{
			Id: option.ID, Data: option.Data})
	}
	return &pb.Question{
		Id:             q.ID,
		Condition:      q.Condition,
		QuestionOption: pbOptions,
	}, nil
}

// service LogicalInterview {
//		TODO:
//     rpc SendAnswer(Ansver) returns (AnswerResponce);
//     rpc GetSummary(User) returns (Summary);
// }
