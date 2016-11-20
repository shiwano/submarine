package debugger

import (
	"image/color"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
)

var (
	colorRed     = &color.RGBA{0xff, 0x41, 0x36, 0xff}
	colorGreen   = &color.RGBA{0x2e, 0xcc, 0x40, 0xff}
	colorBlue    = &color.RGBA{0x00, 0x74, 0xd9, 0xff}
	colorYellow  = &color.RGBA{0xff, 0xdc, 0x00, 0xff}
	colorMaroon  = &color.RGBA{0x85, 0x14, 0x4b, 0xff}
	colorOlive   = &color.RGBA{0x3d, 0x99, 0x70, 0xff}
	colorNavy    = &color.RGBA{0x00, 0x1f, 0x3f, 0xff}
	colorOrange  = &color.RGBA{0xff, 0x85, 0x1b, 0xff}
	colorFuchsia = &color.RGBA{0xf0, 0x12, 0xbe, 0xff}
	colorAqua    = &color.RGBA{0x7f, 0xdb, 0xff, 0xff}
	colorLime    = &color.RGBA{0x01, 0xff, 0x70, 0xff}
	colorPurple  = &color.RGBA{0xb1, 0x0d, 0xc9, 0xff}
	colorTeal    = &color.RGBA{0x39, 0xcc, 0xcc, 0xff}
	colorSilver  = &color.RGBA{0xdd, 0xdd, 0xdd, 0xff}
	colorGray    = &color.RGBA{0xaa, 0xaa, 0xaa, 0xff}
	colorBlack   = &color.RGBA{0x11, 0x11, 0x11, 0xff}
)

func colorByLayer(layer navmesh.LayerMask) (color.Color, navmesh.LayerMask) {
	if layer.Has(navmesh.Layer01) {
		return colorRed, navmesh.Layer01
	} else if layer.Has(navmesh.Layer02) {
		return colorGreen, navmesh.Layer02
	} else if layer.Has(navmesh.Layer03) {
		return colorBlue, navmesh.Layer03
	} else if layer.Has(navmesh.Layer04) {
		return colorYellow, navmesh.Layer04
	} else if layer.Has(navmesh.Layer05) {
		return colorMaroon, navmesh.Layer05
	} else if layer.Has(navmesh.Layer06) {
		return colorOlive, navmesh.Layer06
	} else if layer.Has(navmesh.Layer07) {
		return colorNavy, navmesh.Layer07
	} else if layer.Has(navmesh.Layer08) {
		return colorOrange, navmesh.Layer08
	} else if layer.Has(navmesh.Layer09) {
		return colorFuchsia, navmesh.Layer09
	} else if layer.Has(navmesh.Layer10) {
		return colorAqua, navmesh.Layer10
	} else if layer.Has(navmesh.Layer11) {
		return colorLime, navmesh.Layer11
	} else if layer.Has(navmesh.Layer12) {
		return colorPurple, navmesh.Layer12
	} else if layer.Has(navmesh.Layer13) {
		return colorTeal, navmesh.Layer13
	} else if layer.Has(navmesh.Layer14) {
		return colorSilver, navmesh.Layer14
	} else if layer.Has(navmesh.Layer15) {
		return colorGray, navmesh.Layer15
	}
	return colorBlack, 0
}
