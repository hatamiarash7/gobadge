package dashboard

import (
	"github.com/hatamiarash7/gobadge/svg"
)

type View struct {
	Canvas *svg.Canvas
	Badge  Badge
}

func (view *View) Draw() {
	view.Canvas.Open(float64(badgeW), float64(badgeH))
	view.draw()
	view.Canvas.Close()
}
