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
	screenRect      image.Rectangle
	meshImage       *image.RGBA
	scaledMeshImage *image.RGBA
	objectsImage    *image.RGBA
}

func newNavMeshView(screenRect image.Rectangle) *navMeshView {
	return &navMeshView{
		screenRect: screenRect,
	}
}

func (nm *navMeshView) setScreenRect(screenRect image.Rectangle) {
	nm.screenRect = screenRect
	nm.scaledMeshImage = nm.resizeRGBA(nm.meshImage, nm.screenRect)
	nm.objectsImage = nm.resizeRGBA(nm.objectsImage, nm.screenRect)
}

func (nm *navMeshView) draw(navMesh *navmesh.NavMesh) {
	if nm.meshImage == nil {
		nm.drawMesh(navMesh.Mesh)
	}
	nm.drawObjects(navMesh.Mesh, navMesh.Objects)
}

func (nm *navMeshView) drawMesh(mesh *navmesh.Mesh) {
	meshRectMaxX := int(math.Ceil(mesh.Rect.Max[0])) - int(math.Floor(mesh.Rect.Min[0]))
	meshRectMaxY := int(math.Ceil(mesh.Rect.Max[1])) - int(math.Floor(mesh.Rect.Min[1]))
	meshRect := image.Rect(0, 0, meshRectMaxX, meshRectMaxY)
	nm.meshImage = image.NewRGBA(meshRect)

	gc := draw2dimg.NewGraphicContext(nm.meshImage)
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

	nm.scaledMeshImage = nm.resizeRGBA(nm.meshImage, nm.screenRect)
}

func (nm *navMeshView) drawObjects(mesh *navmesh.Mesh, objects map[int64]navmesh.Object) {
	if nm.objectsImage == nil {
		nm.objectsImage = image.NewRGBA(nm.meshImage.Bounds())
		nm.objectsImage = nm.resizeRGBA(nm.objectsImage, nm.screenRect)
	} else {
		draw.Draw(nm.objectsImage, nm.objectsImage.Bounds(), image.Transparent, image.ZP, draw.Src)
	}

	objectsImageBounds := nm.objectsImage.Bounds()
	meshImageBounds := nm.meshImage.Bounds()
	scaleX := float64(objectsImageBounds.Max.X) / float64(meshImageBounds.Max.X)
	scaleY := float64(objectsImageBounds.Max.Y) / float64(meshImageBounds.Max.Y)

	gc := draw2dimg.NewGraphicContext(nm.objectsImage)
	gc.SetLineWidth(5)
	gc.Scale(scaleX, scaleY)
	for _, o := range objects {
		layer := o.Layer()
		c1, c1Layer := colorByLayer(layer)
		layer.Clear(c1Layer)
		c2, _ := colorByLayer(layer)
		gc.SetFillColor(c1)
		gc.SetStrokeColor(c2)
		p := o.Position()
		x := p[0] - mesh.Rect.Min[0]
		y := p[1] - mesh.Rect.Min[1]
		gc.BeginPath()
		gc.ArcTo(x, y, o.SizeRadius(), o.SizeRadius(), 0, math.Pi*2)
		gc.FillStroke()
	}
}

func (nm *navMeshView) resizeRGBA(src *image.RGBA, rect image.Rectangle) *image.RGBA {
	var r image.Rectangle
	if rect.Max.X > rect.Max.Y {
		r = image.Rect(0, 0, rect.Max.Y, rect.Max.Y)
	} else {
		r = image.Rect(0, 0, rect.Max.X, rect.Max.X)
	}
	dst := image.NewRGBA(r)
	draw.ApproxBiLinear.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Src, nil)
	return dst
}
