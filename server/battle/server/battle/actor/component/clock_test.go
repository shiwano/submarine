package component

import "time"

type fakeClock struct {
	now time.Time
}

func (c *fakeClock) Now() time.Time { return c.now }
