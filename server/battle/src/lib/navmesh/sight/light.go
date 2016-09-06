package sight

import (
	"github.com/ungerik/go3d/float64/vec2"
	"lib/navmesh"
)

type light struct {
	LitPoints []cellPoint `json:"litPoints"`
}

func newLight(navMesh *navmesh.NavMesh, helper *helper, center *vec2.T) *light {
	l := &light{}
	if !navMesh.ContainsPoint(center) {
		return l
	}

	for lightX := 0.0; lightX <= helper.lightDiameter; lightX += helper.cellSize {
		for lightY := 0.0; lightY <= helper.lightDiameter; lightY += helper.cellSize {
			point := &vec2.T{
				lightX - helper.lightRange + center[0],
				lightY - helper.lightRange + center[1],
			}
			if !navMesh.ContainsPoint(point) {
				continue
			}
			vec := vec2.Sub(point, center)
			if vec.LengthSqr() > float64(helper.lightRangeSqr) {
				continue
			}
			if navMesh.Raycast(center, &vec, navmesh.LayerAll) != nil {
				continue
			}
			cellPoint := helper.cellPointByNavMeshPoint(point)
			l.LitPoints = append(l.LitPoints, cellPoint)
		}
	}
	return l
}

func (l *light) isLighting() bool {
	return len(l.LitPoints) > 0
}
