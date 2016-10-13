package ai

import (
	"math"

	"github.com/ungerik/go3d/float64/vec2"

	"github.com/shiwano/submarine/server/battle/server/battle/util"
)

const rad2Deg = 180 / math.Pi

type navigator struct {
	nextPointIndex int
	path           []vec2.T
	previousPoint  *vec2.T
}

func (n *navigator) isStarted() bool {
	return n.path != nil
}

func (n *navigator) start(path []vec2.T, currentPoint *vec2.T) {
	if len(path) < 2 {
		n.stop()
		return
	}
	n.nextPointIndex = 1
	n.path = path
	n.previousPoint = currentPoint
}

func (n *navigator) stop() {
	n.path = nil
	n.nextPointIndex = 0
	n.previousPoint = nil
}

func (n *navigator) navigate(currentPoint *vec2.T) (bool, float64) {
	if !n.isStarted() {
		return false, 0
	}

	nextPoint := &n.path[n.nextPointIndex]
	previousPoint := n.previousPoint
	n.previousPoint = currentPoint

	vec := vec2.Sub(currentPoint, previousPoint)
	vecToNextPoint := vec2.Sub(nextPoint, previousPoint)
	vecLengthSqr := vec.LengthSqr()
	vecToNextPointLengthSqr := vecToNextPoint.LengthSqr()

	if vecLengthSqr < vecToNextPointLengthSqr {
		return true, n.direction(currentPoint, nextPoint)
	}
	dot := vec2.Dot(&vec, &vecToNextPoint)
	if !util.EqualFloats(dot*dot, vecLengthSqr*vecToNextPointLengthSqr) {
		return true, n.direction(currentPoint, nextPoint)
	}

	if n.nextPointIndex == len(n.path)-1 {
		n.stop()
		return false, 0
	}
	n.nextPointIndex++
	return true, n.direction(currentPoint, &n.path[n.nextPointIndex])
}

func (n *navigator) direction(from, to *vec2.T) float64 {
	vec := vec2.Sub(to, from)
	degree := vec.Angle() * rad2Deg
	if degree < 0 {
		degree = 360 + degree
	}
	return degree
}
