package dashboard

import (
	"fmt"
	"strings"

	"github.com/hatamiarash7/gobadge/svg"
	geo "github.com/hatamiarash7/gobadge/svg/geometry"
	"github.com/hatamiarash7/gobadge/svg/style"
)

func (view *View) draw() {
	badge := view.Badge
	origin := geo.Coordinate{X: 0, Y: 0}
	badge.draw(view.Canvas, origin)
}

func (badge Badge) draw(canvas *svg.Canvas, origin geo.Coordinate) error {
	gradient := stockBadgeGradient("#000000", 0.1, "#BBBBBB", 0.1)
	filter := stockTextFilter()

	labelText := stockBadgeLabel(badge.Label)
	tagText := stockBadgeTag(badge.Tag)
	path1 := stockBadgeBlock()
	path2 := stockBadgeColor(colormap(badge.Color))
	rect := stockBadgeGradientRect()
	labelStyle := style.Text{ID: badge.Label, Filter: filter, Text: labelText}
	tagStyle := style.Text{ID: badge.Tag, Filter: filter, Text: tagText}
	translation := fmt.Sprintf("translate(%f,%f)", origin.X, origin.Y)

	output := strings.Join(
		[]string{
			`<g id="` + badge.Label + `" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd" transform="` + translation + `">`,
			`<defs>`,
			gradient.Vectorize(),
			filter.Vectorize(),
			labelText.Vectorize(),
			tagText.Vectorize(),
			`</defs>`,
			`<g>`,
			path1.Vectorize(),
			path2.Vectorize(),
			rect.Vectorize(),
			labelStyle.Vectorize(),
			tagStyle.Vectorize(),
			`</g>`,
			`</g>`,
		}, "\n")

	_, err := fmt.Fprint(canvas.Writer, output)

	return err
}
