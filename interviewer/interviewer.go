package interviewer

import (
	interview "github.com/alex-mj/logical-interview/service"
)

type Interviewer struct {
	DB Store
}

type Store interface {
	Load()
	Store()
}

func (Interviewer) GetNextQuestion(uid interview.UserID) interview.Question {
	return interview.Question{
		ID:        "WE3WR25WERNYW",
		Condition: "Если тряхнуть бурдылькой, то начнется стрельба. Бурдылькой тряхнули.",
		QuestionOptions: []interview.QuestionOption{
			{ID: "1Q1", Data: "стрельба уже началась;"},
			{ID: "1Q2", Data: "стрельба начнется когда-нибудь;"},
			{ID: "1Q3", Data: "стрельба начнется когда-нибудь или уже началась;"},
		},
	}
}
