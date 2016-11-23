package debugger

import (
	"image"
	"image/color"
	"math"

	"github.com/llgcode/draw2d/draw2dimg"
	"golang.org/x/image/draw"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
)

type navMeshView struct {
	screenRect    image.Rectangle
	meshImage     *image.RGBA
	meshImageBase *image.RGBA
	objectsImage  *image.RGBA
}

func newNavMeshView(screenRect image.Rectangle) *navMeshView {
	return &navMeshView{
		screenRect: screenRect,
	}
}

func (nm *navMeshView) setScreenRect(screenRect image.Rectangle) {
	nm.screenRect = screenRect
	nm.meshImage = resizeRGBA(nm.meshImageBase, nm.screenRect)
	nm.objectsImage = resizeRGBA(nm.objectsImage, nm.screenRect)
}

func (nm *navMeshView) draw(navMesh *navmesh.NavMesh) {
	if nm.meshImageBase == nil {
		nm.drawMesh(navMesh.Mesh)
	}
	if nm.objectsImage == nil {
		nm.objectsImage = image.NewRGBA(nm.meshImageBase.Bounds())
		nm.objectsImage = resizeRGBA(nm.objectsImage, nm.screenRect)
	}
	nm.drawObjects(navMesh.Mesh, navMesh.Objects)
}

func (nm *navMeshView) drawMesh(mesh *navmesh.Mesh) {
	meshRectMaxX := int(math.Ceil(mesh.Rect.Max[0])) - int(math.Floor(mesh.Rect.Min[0]))
	meshRectMaxY := int(math.Ceil(mesh.Rect.Max[1])) - int(math.Floor(mesh.Rect.Min[1]))
	meshRect := image.Rect(0, 0, meshRectMaxX, meshRectMaxY)
	nm.meshImageBase = image.NewRGBA(meshRect)

	gc := draw2dimg.NewGraphicContext(nm.meshImageBase)
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
	nm.meshImage = resizeRGBA(nm.meshImageBase, nm.screenRect)
}

func (nm *navMeshView) drawObjects(mesh *navmesh.Mesh, objects map[int64]navmesh.Object) {
	objectsImageBounds := nm.objectsImage.Bounds()
	meshImageBounds := nm.meshImageBase.Bounds()
	scaleX := float64(objectsImageBounds.Max.X) / float64(meshImageBounds.Max.X)
	scaleY := float64(objectsImageBounds.Max.Y) / float64(meshImageBounds.Max.Y)

	draw.Draw(nm.objectsImage, nm.objectsImage.Bounds(), image.Transparent, image.ZP, draw.Src)
	gc := draw2dimg.NewGraphicContext(nm.objectsImage)
	gc.Scale(scaleX, scaleY)
	objectSlice := newObjectSlice(objects)
	objectSlice.sort()
	for _, o := range objectSlice {
		nm.drawObject(gc, o, mesh)
	}
}

func (nm *navMeshView) drawObject(gc *draw2dimg.GraphicContext, o navmesh.Object, mesh *navmesh.Mesh) {
	colors := colorsByLayer(o.Layer(), 2)

	p := o.Position()
	x := p[0] - mesh.Rect.Min[0]
	y := p[1] - mesh.Rect.Min[1]
	lineWidth := o.SizeRadius() / 2
	r := o.SizeRadius() - lineWidth/2

	gc.SetLineWidth(lineWidth)
	gc.SetFillColor(colors[0])
	gc.SetStrokeColor(colors[1])
	gc.BeginPath()
	gc.ArcTo(x, y, r, r, 0, math.Pi*2)
	gc.FillStroke()
}
