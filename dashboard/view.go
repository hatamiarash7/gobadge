package dashboard

import (
	"github.com/hatamiarash7/gobadge/svg"
)

type View struct {
	Canvas *svg.Canvas
	Badge  Badge
}

func (view *View) Draw() {
	view.Canvas.Open(view.Badge.endWidth(), float64(20))
	view.draw()
	view.Canvas.Close()
}
