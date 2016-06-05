package context

import (
	"github.com/ungerik/go3d/float64/vec2"
)

// User represents an user in the battle.
type User struct {
	ID            int64
	StartPosition *vec2.T
}
