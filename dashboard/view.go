package dashboard

import (
	"github.com/hatamiarash7/gobadge/svg"
)

type View struct {
	Canvas *svg.Canvas
	Badge  Badge
}

func (view *View) Draw() {
	labelWidth := float64(len(view.Badge.Label)) * 6.5
	tagWidth := float64(len(view.Badge.Tag)) * 6.5
	end := labelWidth + tagWidth + 13

	view.Canvas.Open(end, float64(20))
	view.draw()
	view.Canvas.Close()
}
