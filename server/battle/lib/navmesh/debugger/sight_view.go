package debugger

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
	"golang.org/x/image/draw"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	"github.com/shiwano/submarine/server/battle/lib/navmesh/sight"
)

type sightView struct {
	screenRect image.Rectangle
	sightImage *image.RGBA
}

func newSightView(screenRect image.Rectangle) *sightView {
	return &sightView{
		screenRect: screenRect,
	}
}

func (sv *sightView) setScreenRect(screenRect image.Rectangle) {
	sv.screenRect = screenRect
	sv.sightImage = resizeRGBA(sv.sightImage, sv.screenRect)
}

func (sv *sightView) draw(s *sight.Sight, mesh *navmesh.Mesh) {
	meshImageRect := imageRectFromVec2Rect(mesh.Rect)

	if sv.sightImage == nil {
		sv.sightImage = image.NewRGBA(meshImageRect)
		sv.sightImage = resizeRGBA(sv.sightImage, sv.screenRect)
	}

	scaleX, scaleY := scaleValues(sv.sightImage.Bounds(), meshImageRect)

	draw.Draw(sv.sightImage, sv.sightImage.Bounds(), image.Transparent, image.ZP, draw.Src)
	gc := draw2dimg.NewGraphicContext(sv.sightImage)
	gc.Scale(scaleX, scaleY)
	gc.SetFillColor(color.RGBA{0x00, 0x00, 0x00, 0x66})

	for _, p := range s.LitPoints() {
		x := p[0] - mesh.Rect.Min[0]
		y := p[1] - mesh.Rect.Min[1]
		gc.MoveTo(x, y)
		gc.LineTo(x+s.CellSize(), y)
		gc.LineTo(x+s.CellSize(), y-s.CellSize())
		gc.LineTo(x, y-s.CellSize())
		gc.LineTo(x, y)
		gc.FillStroke()
		gc.Fill()
		gc.Close()
	}
}
