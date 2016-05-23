package actor

import (
	"app/battle/context"
	"app/typhenapi/type/submarine/battle"
	"github.com/ungerik/go3d/float64/vec2"
	"lib/currentmillis"
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
		context:            context,
		initialPosition:    position,
		normalizedVelocity: &vec2.T{1, 0},
		changedAt:          context.Now,
		accelerator: &accelerator{
			context:    context,
			maxSpeed:   maxSpeed,
			duration:   duration,
			isShutdown: true,
		},
	}
	return m
}

func (m *motor) toAPIType(actorID int64) *battle.Movement {
	position := m.position()
	return &battle.Movement{
		ActorId:     actorID,
		Position:    &battle.Point{X: position[0], Y: position[1]},
		Direction:   m.direction,
		MovedAt:     currentmillis.Millis(m.changedAt),
		Accelerator: m.accelerator.toAPIType(),
	}
}

func (m *motor) accelerate() {
	m.initialPosition = m.position()
	m.initialSpeed = m.accelerator.speed()
	m.accelerator.refresh(true, false)
	m.changedAt = m.context.Now
}

func (m *motor) brake() {
	m.initialPosition = m.position()
	m.initialSpeed = m.accelerator.speed()
	m.accelerator.refresh(false, false)
	m.changedAt = m.context.Now
}

func (m *motor) turn(direction float64) {
	m.initialPosition = m.position()
	m.initialSpeed = m.accelerator.speed()
	m.accelerator.refresh(m.accelerator.isAccelerating, false)
	m.direction = direction
	m.normalizedVelocity = &vec2.T{
		math.Cos(direction * deg2Rad),
		math.Sin(direction * deg2Rad),
	}
	m.changedAt = m.context.Now
}

func (m *motor) idle(idlingPosition *vec2.T) {
	m.initialPosition = idlingPosition
	m.initialSpeed = m.accelerator.speed()
	m.accelerator.refresh(m.accelerator.isAccelerating, true)
	m.changedAt = m.context.Now
}

func (m *motor) position() *vec2.T {
	if m.accelerator.isIdling || m.accelerator.isShutdown {
		return &(*m.initialPosition)
	}

	var t1, t2 float64
	if m.context.Now.After(m.accelerator.reachedMaxSpeedAt) {
		t1 = m.accelerator.reachedMaxSpeedAt.Sub(m.changedAt).Seconds()
		t2 = m.context.Now.Sub(m.accelerator.reachedMaxSpeedAt).Seconds()
	} else {
		t1 = m.context.Now.Sub(m.changedAt).Seconds()
	}

	p := *m.initialPosition
	v := m.initialSpeed * t1
	a := m.accelerator.acceleration() * math.Pow(t1, 2) / 2
	d1 := m.normalizedVelocity.Scaled(v + a)
	if m.accelerator.isAccelerating {
		d2 := m.normalizedVelocity.Scaled(m.accelerator.maxSpeed * t2)
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
	isIdling          bool
	isAccelerating    bool
	changedAt         time.Time
	reachedMaxSpeedAt time.Time
}

func (a *accelerator) toAPIType() *battle.Accelerator {
	if a.isIdling || a.isShutdown {
		return nil
	}
	return &battle.Accelerator{
		MaxSpeed:       a.maxSpeed,
		Duration:       currentmillis.DurationMillis(a.duration),
		StartRate:      a.startRate,
		IsAccelerating: a.isAccelerating,
	}
}

func (a *accelerator) refresh(isAccelerating bool, isIdling bool) {
	a.startRate = a.rate()
	a.isShutdown = false
	a.isIdling = isIdling
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
