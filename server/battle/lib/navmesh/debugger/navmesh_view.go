package debugger

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/llgcode/draw2d/draw2dimg"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
)

type navMeshView struct {
	navMesh       *navmesh.NavMesh
	meshImage     *image.RGBA
	meshImageBase *image.RGBA
}

func newNavMeshView(navMesh *navmesh.NavMesh) *navMeshView {
	return &navMeshView{
		navMesh: navMesh,
	}
}

func (nv *navMeshView) draw(im *image.RGBA) {
	if nv.meshImageBase == nil {
		nv.meshImageBase = nv.createMeshImageBase()
	}
	if nv.meshImage == nil || !nv.meshImage.Bounds().Eq(imageRectForResize(im.Bounds())) {
		nv.meshImage = resizeRGBA(nv.meshImageBase, im.Bounds())
	}
	draw.Draw(im, im.Bounds(), nv.meshImage, image.ZP, draw.Src)
	nv.drawObjects(im)
}

func (nv *navMeshView) createMeshImageBase() *image.RGBA {
	mesh := nv.navMesh.Mesh
	meshImageRect := imageRectFromVec2Rect(mesh.Rect)
	meshImageBase := image.NewRGBA(meshImageRect)

	gc := draw2dimg.NewGraphicContext(meshImageBase)
	gc.SetFillColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	gc.SetLineWidth(1)

	for _, t := range mesh.Triangles {
		gc.MoveTo(t.Vertices[0][0]-mesh.Rect.Min[0], t.Vertices[0][1]-mesh.Rect.Min[1])
		gc.LineTo(t.Vertices[1][0]-mesh.Rect.Min[0], t.Vertices[1][1]-mesh.Rect.Min[1])
		gc.LineTo(t.Vertices[2][0]-mesh.Rect.Min[0], t.Vertices[2][1]-mesh.Rect.Min[1])
		gc.LineTo(t.Vertices[0][0]-mesh.Rect.Min[0], t.Vertices[0][1]-mesh.Rect.Min[1])
		gc.FillStroke()
		gc.Close()
	}
	return meshImageBase
}

func (nv *navMeshView) drawObjects(im *image.RGBA) {
	scaleX, scaleY := scaleValues(imageRectForResize(im.Bounds()), nv.meshImageBase.Bounds())

	gc := draw2dimg.NewGraphicContext(im)
	gc.Scale(scaleX, scaleY)
	objectSlice := newObjectSlice(nv.navMesh.Objects)
	objectSlice.sort()

	for _, o := range objectSlice {
		nv.drawObject(gc, o)
	}
}

func (nv *navMeshView) drawObject(gc *draw2dimg.GraphicContext, o navmesh.Object) {
	colors := colorsByLayer(o.Layer(), 2)

	p := o.Position()
	x := p[0] - nv.navMesh.Mesh.Rect.Min[0]
	y := p[1] - nv.navMesh.Mesh.Rect.Min[1]
	lineWidth := o.SizeRadius() / 2
	r := o.SizeRadius() - lineWidth/2

	gc.SetLineWidth(lineWidth)
	gc.SetFillColor(colors[0])
	gc.SetStrokeColor(colors[1])
	gc.BeginPath()
	gc.ArcTo(x, y, r, r, 0, math.Pi*2)
	gc.FillStroke()
}
