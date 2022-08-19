package dashboard

import (
	"fmt"
	"strings"

	"github.com/hatamiarash7/gobadge/svg"
	"github.com/hatamiarash7/gobadge/svg/style"
)

func (view *View) draw() {
	badge := view.Badge
	badge.draw(view.Canvas)
}

func (badge Badge) draw(canvas *svg.Canvas) error {
	gradient := stockBadgeGradient("#000000", 0.1, "#BBBBBB", 0.1)
	filter := stockTextFilter()

	labelText := stockBadgeLabel(badge.Label)
	tagText := stockBadgeTag(badge.Tag, badge.labelWidth())
	path1 := stockBadgeBlock(badge.labelWidth())
	path2 := stockBadgeColor(colormap(badge.Color), badge.labelWidth(), badge.endWidth())
	rect := stockBadgeGradientRect(badge.endWidth())
	labelStyle := style.Text{ID: badge.Label, Filter: filter, Text: labelText}
	tagStyle := style.Text{ID: badge.Tag, Filter: filter, Text: tagText}

	output := strings.Join(
		[]string{
			`<g id="` + badge.Label + `" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">`,
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
