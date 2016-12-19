package component

import (
	"math"
	"time"

	"github.com/shiwano/submarine/server/battle/lib/currentmillis"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"github.com/shiwano/submarine/server/battle/server/battle/context"

	"github.com/ungerik/go3d/float64/vec2"
)

const (
	rad2Deg = 180 / math.Pi
	deg2Rad = math.Pi / 180
)

// Motor represents a motor for actor moving.
type Motor struct {
	ctx                *context.Context
	accelerator        *accelerator
	initialPosition    *vec2.T
	initialSpeed       float64
	direction          float64
	normalizedVelocity *vec2.T
	changedAt          time.Time
}

// NewMotor creates a motor.
func NewMotor(ctx *context.Context, position *vec2.T, direction float64,
	maxSpeed float64, duration time.Duration) *Motor {
	return &Motor{
		ctx:             ctx,
		initialPosition: position,
		direction:       direction,
		normalizedVelocity: &vec2.T{
			math.Cos(direction * deg2Rad),
			math.Sin(direction * deg2Rad),
		},
		changedAt: ctx.Now,
		accelerator: &accelerator{
			ctx:        ctx,
			maxSpeed:   maxSpeed,
			duration:   duration,
			isShutdown: true,
		},
	}
}

// Direction returns the direction of the motor moving.
func (m *Motor) Direction() float64 { return m.direction }

// IsAccelerating determines whether the motor is accelerating.
func (m *Motor) IsAccelerating() bool { return m.accelerator.isAccelerating }

// NormalizedVelocity returns the normalized velocity of the motor moving.
func (m *Motor) NormalizedVelocity() vec2.T { return *m.normalizedVelocity }

// ToAPIType creates a movement message from the motor.
func (m *Motor) ToAPIType(actorID int64) *battleAPI.Movement {
	return &battleAPI.Movement{
		ActorId:     actorID,
		Position:    &battleAPI.Point{X: m.initialPosition[0], Y: m.initialPosition[1]},
		Direction:   m.direction,
		MovedAt:     currentmillis.Millis(m.changedAt),
		Accelerator: m.accelerator.toAPIType(),
	}
}

// Accelerate the motor.
func (m *Motor) Accelerate(position *vec2.T) {
	m.initialPosition = position
	m.initialSpeed = m.accelerator.speed()
	m.accelerator.refresh(true, false)
	m.changedAt = m.ctx.Now
}

// Brake the motor.
func (m *Motor) Brake(position *vec2.T) {
	m.initialPosition = m.Position()
	m.initialSpeed = m.accelerator.speed()
	m.accelerator.refresh(false, false)
	m.changedAt = m.ctx.Now
}

// Turn the direction of the motor.
func (m *Motor) Turn(position *vec2.T, direction float64) {
	m.initialPosition = position
	m.initialSpeed = m.accelerator.speed()
	m.accelerator.refresh(m.accelerator.isAccelerating, false)
	m.direction = direction
	m.normalizedVelocity = &vec2.T{
		math.Cos(direction * deg2Rad),
		math.Sin(direction * deg2Rad),
	}
	m.changedAt = m.ctx.Now
}

// Idle the motor.
func (m *Motor) Idle(position *vec2.T) {
	m.initialPosition = position
	m.initialSpeed = m.accelerator.speed()
	m.accelerator.refresh(m.accelerator.isAccelerating, true)
	m.changedAt = m.ctx.Now
}

// Position returns the current position.
func (m *Motor) Position() *vec2.T {
	if m.accelerator.isIdling || m.accelerator.isShutdown {
		return &(*m.initialPosition)
	}

	var t1, t2 float64
	if m.ctx.Now.After(m.accelerator.reachedMaxSpeedAt) {
		t1 = m.accelerator.reachedMaxSpeedAt.Sub(m.changedAt).Seconds()
		t2 = m.ctx.Now.Sub(m.accelerator.reachedMaxSpeedAt).Seconds()
	} else {
		t1 = m.ctx.Now.Sub(m.changedAt).Seconds()
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
	ctx               *context.Context
	maxSpeed          float64
	duration          time.Duration
	startRate         float64
	isShutdown        bool
	isIdling          bool
	isAccelerating    bool
	changedAt         time.Time
	reachedMaxSpeedAt time.Time
}

func (a *accelerator) toAPIType() *battleAPI.Accelerator {
	if a.isIdling || a.isShutdown {
		return nil
	}
	return &battleAPI.Accelerator{
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
	a.changedAt = a.ctx.Now
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
	rate := a.ctx.Now.Sub(a.changedAt).Seconds() / a.duration.Seconds()
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
