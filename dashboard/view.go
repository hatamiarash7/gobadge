package dashboard

import (
	"github.com/hatamiarash7/gobadge/svg"
)

// View is the main struct for our view
type View struct {
	Canvas *svg.Canvas
	Badge  Badge
}

// Draw will draw the badge
func (view *View) Draw() {
	err := view.Canvas.Open(view.Badge.endWidth(), float64(20))

	if err != nil {
		panic(err)
	}

	view.draw()
	view.Canvas.Close()
}
