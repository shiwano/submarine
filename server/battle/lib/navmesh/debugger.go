package navmesh

// Debugger is the Debugger interface of the debugger package.
type Debugger interface {
	UpdateNavMesh(navMesh *NavMesh)
}
