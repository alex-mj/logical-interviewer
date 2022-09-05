package interviewer

type Interviewer struct {
	DB Store
}

type Store interface {
	Load()
	Store()
}
