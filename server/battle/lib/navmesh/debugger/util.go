package debugger

import (
	"image"
	"math"

	"golang.org/x/image/draw"

	"github.com/ungerik/go3d/float64/vec2"
)

func imageRectFromVec2Rect(r *vec2.Rect) image.Rectangle {
	maxX := int(math.Ceil(r.Max[0])) - int(math.Floor(r.Min[0]))
	maxY := int(math.Ceil(r.Max[1])) - int(math.Floor(r.Min[1]))
	return image.Rect(0, 0, maxX, maxY)
}

func scaleValues(a image.Rectangle, b image.Rectangle) (x float64, y float64) {
	x = float64(a.Min.X+a.Max.X) / float64(b.Min.Y+b.Max.X)
	y = float64(a.Min.Y+a.Max.Y) / float64(b.Min.Y+b.Max.Y)
	return
}

func resizeRGBA(src *image.RGBA, rect image.Rectangle) *image.RGBA {
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
