package navmesh

// LayerMask represents a layer of the navmesh.
type LayerMask uint16

// Layer constants.
const (
	Layer0 LayerMask = 1 << iota
	Layer1
	Layer2
	Layer3
	Layer4
	Layer5
	Layer6
	Layer7
	Layer8
	Layer9
	Layer10
	Layer11
	Layer12
	Layer13
	Layer14
	Layer15
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
