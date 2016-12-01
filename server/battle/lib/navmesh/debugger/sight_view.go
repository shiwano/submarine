package debugger

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	"github.com/shiwano/submarine/server/battle/lib/navmesh/sight"
)

type sightView struct {
	sight *sight.Sight
	mesh  *navmesh.Mesh
}

func newSightView(sight *sight.Sight, mesh *navmesh.Mesh) *sightView {
	return &sightView{
		sight: sight,
		mesh:  mesh,
	}
}

func (sv *sightView) draw(im *image.RGBA) {
	meshImageRect := imageRectFromVec2Rect(sv.mesh.Rect)
	scaleX, scaleY := scaleValues(imageRectForResize(im.Bounds()), meshImageRect)

	gc := draw2dimg.NewGraphicContext(im)
	gc.Scale(scaleX, scaleY)
	gc.SetFillColor(color.RGBA{0x00, 0x00, 0x00, 0x66})

	for _, p := range sv.sight.LitPoints() {
		x := p[0] - sv.mesh.Rect.Min[0]
		y := p[1] - sv.mesh.Rect.Min[1]
		gc.MoveTo(x, y)
		gc.LineTo(x+sv.sight.CellSize(), y)
		gc.LineTo(x+sv.sight.CellSize(), y-sv.sight.CellSize())
		gc.LineTo(x, y-sv.sight.CellSize())
		gc.LineTo(x, y)
		gc.FillStroke()
		gc.Fill()
		gc.Close()
	}
}
