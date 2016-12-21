package component

// Visibility represents visibility of an actor, it manages visibility in duplicate.
type Visibility struct {
	count         int
	ChangeHandler func()
}

// IsVisible determines whether the actor is visible.
func (v *Visibility) IsVisible() bool {
	return v.count > 0
}

// Set the actor visibility.
func (v *Visibility) Set(isVisible bool) {
	previousVisibility := v.IsVisible()

	if isVisible {
		v.count++
	} else if v.count > 0 {
		v.count--
	}

	if v.ChangeHandler != nil && previousVisibility != v.IsVisible() {
		v.ChangeHandler()
	}
}
