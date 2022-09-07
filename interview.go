package domain

import "time"

type UserID string

type Question struct {
	ID              string
	Condition       string
	QuestionOptions []QuestionOption
}

type QuestionOption struct {
	ID             string
	Data           string
	CustomerChoice bool
	Correct        bool
}

type Ansver struct {
	UserID     UserID
	QuestionID string
	AnsverID   string
}

type AnsverResponse struct {
	Ansvered int32
	Goal     int32
	TimeLeft time.Duration
}

type LastResult struct {
	Scores          int32
	TimeSpent       time.Duration
	TimeTest        time.Time
	NewAttempInXDay int32
}

type Interviewer interface {
	GetNextQuestion(UserID) Question
}
