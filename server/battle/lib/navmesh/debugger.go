package navmesh

// Debugger is the Debugger interface. See also debugger package.
type Debugger interface {
	Update(navMesh *NavMesh)
}
