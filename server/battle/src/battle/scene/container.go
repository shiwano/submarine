package scene

import battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"

type container struct {
	actors               ActorSlice
	actorsByID           map[int64]Actor
	submarinesByPlayerID map[int64]Actor
	players              PlayerSlice
	userPlayersByTeam    PlayersByTeam
}

func newContainer() *container {
	c := &container{
		actorsByID:           make(map[int64]Actor),
		submarinesByPlayerID: make(map[int64]Actor),
	}
	return c
}

func (c *container) addActor(actor Actor) {
	c.actorsByID[actor.ID()] = actor
	c.actors = append(c.actors, actor)
	if actor.Type() == battleAPI.ActorType_Submarine {
		c.submarinesByPlayerID[actor.Player().ID] = actor
		c.players = append(c.players, actor.Player())
	}
	c.userPlayersByTeam = c.players.Where(func(p *Player) bool {
		return p.IsUser
	}).GroupByTeam()
}

func (c *container) removeActor(rawActor Actor) Actor {
	actor := c.actorsByID[rawActor.ID()]
	if actor == nil {
		return nil
	}

	delete(c.actorsByID, actor.ID())
	actors := c.actors[:0]
	for _, a := range c.actors {
		if a != actor {
			actors = append(actors, a)
		}
	}
	c.actors = actors
	if actor.Type() == battleAPI.ActorType_Submarine {
		delete(c.submarinesByPlayerID, actor.Player().ID)

		players := c.players[:0]
		for _, p := range c.players {
			if p != actor.Player() {
				players = append(players, p)
			}
		}
		c.players = players
	}
	c.userPlayersByTeam = c.players.Where(func(p *Player) bool {
		return p.IsUser
	}).GroupByTeam()
	return actor
}
