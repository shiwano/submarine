package debugger

import (
	"image"
	"log"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

// Main is called by the program's main function to run the debugger.
func Main(callback func(*Debugger)) {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(nil)
		if err != nil {
			log.Fatal(err)
		}
		defer w.Release()

		var b screen.Buffer
		defer func() {
			if b != nil {
				b.Release()
			}
		}()

		debugger := newDebugger(w)

		var hasPublished bool
		for {
			e := w.NextEvent()

			switch e := e.(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			case key.Event:
				if e.Code == key.CodeEscape {
					return
				}
			case paint.Event:
				w.Upload(image.Point{}, b, b.Bounds())
				w.Publish()
				if !hasPublished {
					hasPublished = true
					if callback != nil {
						go callback(debugger)
					}
				}
			case size.Event:
				if e.WidthPx == 0 && e.HeightPx == 0 {
					return
				}
				b, err = s.NewBuffer(e.Size())
				if err != nil {
					log.Fatal(err)
				}
				debugger.setScreen(b.RGBA())
				debugger.render()
			case error:
				log.Fatal(err)
			}
		}
	})
}
