package battle

import (
	"time"

	"github.com/shiwano/submarine/server/battle/src/battle/context"
)

type judge struct {
	context   context.FullContext
	timeLimit time.Duration
}

func newJudge(context context.FullContext, timeLimit time.Duration) *judge {
	return &judge{
		context:   context,
		timeLimit: timeLimit,
	}
}

func (j judge) isBattleFinished() bool {
	return j.context.ElapsedTime() >= j.timeLimit || len(j.context.Players()) == 1
}

func (j judge) winner() *context.Player {
	users := j.context.Players()
	if len(users) != 1 {
		return nil
	}
	return users[0]
}
