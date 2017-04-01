package component

import "time"

type clock interface {
	Now() time.Time
}
