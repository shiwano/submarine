package sight

import (
	"github.com/ungerik/go3d/float64/vec2"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
)

type light struct {
	LitPoints []cellPoint `json:"litPoints"`
}

func newLight(navMesh *navmesh.NavMesh, helper *helper, center *vec2.T) *light {
	l := &light{}
	if !navMesh.ContainsPoint(center) {
		return l
	}
	lightRangeSqr := helper.LightRange * helper.LightRange
	lightDiameter := helper.LightRange*2 + 1

	for lightX := 0.0; lightX <= lightDiameter; lightX += helper.CellSize {
		for lightY := 0.0; lightY <= lightDiameter; lightY += helper.CellSize {
			point := &vec2.T{
				lightX - helper.LightRange + center[0],
				lightY - helper.LightRange + center[1],
			}
			if !navMesh.ContainsPoint(point) {
				continue
			}
			vec := vec2.Sub(point, center)
			if vec.LengthSqr() > float64(lightRangeSqr) {
				continue
			}
			var cellPoint cellPoint
			if hitInfo := navMesh.Raycast(center, &vec, navmesh.LayerAll); hitInfo != nil {
				cellPoint = helper.cellPointByNavMeshPoint(&hitInfo.Point)
			} else {
				cellPoint = helper.cellPointByNavMeshPoint(point)
			}
			l.LitPoints = append(l.LitPoints, cellPoint)
		}
	}
	return l
}

func (l *light) isLighting() bool {
	return len(l.LitPoints) > 0
}
