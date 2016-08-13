package context

// AI represents a battle artificial intelligence.
type AI interface {
	Update(Actor)
}
