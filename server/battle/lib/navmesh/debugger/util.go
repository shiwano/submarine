package debugger

import (
	"image"

	"golang.org/x/image/draw"
)

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
