package battle

import (
	"time"

	"github.com/shiwano/submarine/server/battle/src/battle/scene"
)

type judge struct {
	scene     scene.FullScene
	timeLimit time.Duration
}

func newJudge(scene scene.FullScene, timeLimit time.Duration) *judge {
	return &judge{
		scene:     scene,
		timeLimit: timeLimit,
	}
}

func (j judge) isBattleFinished() bool {
	return j.scene.ElapsedTime() >= j.timeLimit || len(j.scene.Players()) == 1
}

func (j judge) winner() *scene.Player {
	users := j.scene.Players()
	if len(users) != 1 {
		return nil
	}
	return users[0]
}
