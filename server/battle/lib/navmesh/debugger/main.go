package debugger

import (
	"log"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

// Main is called by the program's main function to run the debugger.
// Before executing this, it is preferable to execute runtime.LockOSThread().
// Note the callback will be executed by another goroutine.
func Main(callback func(*Debugger)) {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(nil)
		if err != nil {
			log.Fatal(err)
		}
		defer w.Release()

		d := newDebugger(w)
		defer d.close()

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
				d.uploadScreen(w)
				w.Publish()
				if !hasPublished {
					hasPublished = true
					if callback != nil {
						go callback(d)
					}
				}
			case size.Event:
				if e.WidthPx == 0 && e.HeightPx == 0 {
					return
				}
				b, err := s.NewBuffer(e.Size())
				if err != nil {
					log.Fatal(err)
				}
				d.setScreen(b)
			case error:
				log.Fatal(err)
			}
		}
	})
}
