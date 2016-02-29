package actor

import (
	"app/battle/context"
	"app/currentmillis"
	"app/typhenapi/type/submarine/battle"
	"github.com/ungerik/go3d/float64/vec2"
	"math"
	"time"
)

const (
	rad2Deg = 180 / math.Pi
	deg2Rad = math.Pi / 180
)

type motor struct {
	context            *context.Context
	accelerator        *accelerator
	initialPosition    *vec2.T
	initialSpeed       float64
	direction          float64
	normalizedVelocity *vec2.T
	changedAt          time.Time
}

func newMotor(context *context.Context, position *vec2.T,
	maxSpeed float64, duration time.Duration) *motor {
	m := &motor{
		context: context,
		accelerator: &accelerator{
			context:    context,
			maxSpeed:   maxSpeed,
			duration:   duration,
			isShutdown: true,
		},
	}
	m.turn(position, 0)
	return m
}

func (m *motor) toAPIType(actorID int64) *battle.Movement {
	position := m.position()
	return &battle.Movement{
		ActorId:     actorID,
		Position:    &battle.Point{X: position[0], Y: position[1]},
		Direction:   m.direction,
		MovedAt:     currentmillis.Milliseconds(m.changedAt),
		Accelerator: m.accelerator.toAPIType(),
	}
}

func (m *motor) accelerate(position *vec2.T) {
	m.initialPosition = position
	m.initialSpeed = m.accelerator.speed()
	m.accelerator.switchAccelerating(true)
	m.changedAt = m.context.Now
}

func (m *motor) brake(position *vec2.T) {
	m.initialPosition = position
	m.initialSpeed = m.accelerator.speed()
	m.accelerator.switchAccelerating(false)
	m.changedAt = m.context.Now
}

func (m *motor) turn(position *vec2.T, direction float64) {
	m.initialPosition = position
	m.initialSpeed = m.accelerator.speed()
	m.direction = direction
	m.normalizedVelocity = &vec2.T{
		math.Cos(direction * deg2Rad),
		math.Sin(direction * deg2Rad),
	}
	m.changedAt = m.context.Now
}

func (m *motor) position() *vec2.T {
	var s1, s2 float64
	if m.context.Now.After(m.accelerator.reachedMaxSpeedAt) {
		s1 = m.accelerator.reachedMaxSpeedAt.Sub(m.accelerator.changedAt).Seconds()
		s2 = m.context.Now.Sub(m.accelerator.reachedMaxSpeedAt).Seconds()
	} else {
		s1 = m.context.Now.Sub(m.changedAt).Seconds()
	}

	p := *m.initialPosition
	v := m.initialSpeed * s1
	a := m.accelerator.acceleration() * math.Pow(s1, 2) / 2
	d1 := m.normalizedVelocity.Scaled(v + a)
	if m.accelerator.isAccelerating {
		d2 := m.normalizedVelocity.Scaled(m.accelerator.maxSpeed * s2)
		return p.Add(&d1).Add(&d2)
	}
	return p.Add(&d1)
}

type accelerator struct {
	context           *context.Context
	maxSpeed          float64
	duration          time.Duration
	startRate         float64
	isShutdown        bool
	isAccelerating    bool
	changedAt         time.Time
	reachedMaxSpeedAt time.Time
}

func (a *accelerator) toAPIType() *battle.Accelerator {
	if a.isShutdown {
		return nil
	}
	return &battle.Accelerator{
		MaxSpeed:       a.maxSpeed,
		Duration:       currentmillis.MillisecondsDuration(a.duration),
		StartRate:      a.startRate,
		IsAccelerating: a.isAccelerating,
		ChangedAt:      currentmillis.Milliseconds(a.changedAt),
	}
}

func (a *accelerator) switchAccelerating(isAccelerating bool) {
	a.startRate = a.rate()
	a.isShutdown = false
	a.isAccelerating = isAccelerating
	a.changedAt = a.context.Now
	remainingRate := a.startRate
	if a.isAccelerating {
		remainingRate = 1 - remainingRate
	}
	remainingNanoSeconds := float64(a.duration.Nanoseconds()) * remainingRate
	a.reachedMaxSpeedAt = a.changedAt.Add(time.Duration(remainingNanoSeconds))
}

func (a *accelerator) rate() float64 {
	if a.isShutdown {
		return 0
	}
	rate := a.context.Now.Sub(a.changedAt).Seconds() / a.duration.Seconds()
	if !a.isAccelerating {
		rate = -rate
	}
	rate = a.startRate + rate
	if rate >= 1 {
		return 1
	} else if rate <= 0 {
		return 0
	}
	return rate
}

func (a *accelerator) acceleration() float64 {
	if a.isShutdown {
		return 0
	}
	if !a.isAccelerating {
		return -a.maxSpeed / a.duration.Seconds()
	}
	return a.maxSpeed / a.duration.Seconds()
}

func (a *accelerator) speed() float64 {
	rate := a.rate()
	if rate >= 1 {
		return a.maxSpeed
	} else if rate <= 0 {
		return 0
	}
	return a.maxSpeed * rate
}
