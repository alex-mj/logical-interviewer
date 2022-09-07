package interviewer

import (
	"strconv"
	"testing"
	"time"

	interview "github.com/alex-mj/logical-interview/service"
	"github.com/stretchr/testify/require"
)

func TestOk(t *testing.T) {
	require.Equal(t, 1, 1, "we are in the same galactic")
}

func TestGetNextQuestion(t *testing.T) {

	intwr := Interviewer{}

	q := intwr.GetNextQuestion(
		interview.UserID(
			strconv.Itoa(int(time.Now().UnixMicro())),
		))

	require.Greater(t, len(q.QuestionOptions), 0, "expected some question options")
}
