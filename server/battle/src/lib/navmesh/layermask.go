package navmesh

// LayerMask represents a layer of the navmesh.
type LayerMask uint16

// Layer constants.
const (
	Layer01 LayerMask = 1 << iota
	Layer02
	Layer03
	Layer04
	Layer05
	Layer06
	Layer07
	Layer08
	Layer09
	Layer10
	Layer11
	Layer12
	Layer13
	Layer14
	Layer15
	Layer16
)

// Has determines whether the layer mask has specified layers.
func (l LayerMask) Has(layer LayerMask) bool {
	return l&layer != 0
}

// Set sets specified bit flags on the layer mask.
func (l *LayerMask) Set(layer LayerMask) {
	*l |= layer
}

// Clear clears specified bit flags on the layer mask.
func (l *LayerMask) Clear(layer LayerMask) {
	*l &= ^layer
}
