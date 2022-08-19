package dashboard

import (
	"fmt"
	"strings"

	geo "github.com/hatamiarash7/gobadge/svg/geometry"
	"github.com/hatamiarash7/gobadge/svg/shapes"
	"github.com/hatamiarash7/gobadge/svg/style"
)

func stockBadgeLabel(text string) shapes.Text {
	id := makeTextID(text, "label")
	origin := geo.Coordinate{X: 8, Y: 14}
	return stockText(text, id, origin)
}

func stockBadgeTag(text string, label string) shapes.Text {
	labelWidth := float64(len(label)) * 6.5

	id := makeTextID(text, "tag")
	origin := geo.Coordinate{X: labelWidth + 6, Y: 14}
	return stockText(text, id, origin)
}

func stockBadgeBlock(text string) shapes.Path {
	width := float64(len(text)) * 6.5

	return shapes.Path{
		ID:       "block",
		Fill:     "#000",
		FillRule: "nonzero",
		D:        "M 3 0 L " + fmt.Sprintf("%f", width) + " 0 L " + fmt.Sprintf("%f", width) + " 20 L 3 20 C 1.3431 20 0 18.6569 0 17 L 0 3 C 0 1.3431 1.3431 0 3 0",
	}
}

func stockBadgeColor(fill string, label string, tag string) shapes.Path {
	labelWidth := float64(len(label)) * 6.5
	tagWidth := float64(len(tag)) * 6.5
	end := labelWidth + tagWidth + 13

	return shapes.Path{
		ID:       "color",
		Fill:     fill,
		FillRule: "nonzero",
		D:        "M " + fmt.Sprintf("%f", labelWidth) + " 0 L " + fmt.Sprintf("%f", end-3) + " 0 C " + fmt.Sprintf("%f", end-2+0.6569) + " 0 " + fmt.Sprintf("%f", end) + " 2 " + fmt.Sprintf("%f", end) + " 3 L " + fmt.Sprintf("%f", end) + " 17 C " + fmt.Sprintf("%f", end) + " 18.6569 " + fmt.Sprintf("%f", end-2+0.6569) + " 20 " + fmt.Sprintf("%f", end-3) + " 20 L " + fmt.Sprintf("%f", labelWidth) + " 20",
	}
}

func stockBadgeGradientRect(label string, tag string) shapes.Rect {
	labelWidth := float64(len(label)) * 6.5
	tagWidth := float64(len(tag)) * 6.5
	end := labelWidth + tagWidth + 13

	return shapes.Rect{
		ID:       "gradient",
		Fill:     "url(#texture)",
		FillRule: "nonzero",
		Origin:   geo.Coordinate{X: 0, Y: 0},
		Radius:   geo.Radius{X: 3, Y: 3},
		Frame:    geo.Frame{Width: end, Height: 20},
	}
}

func stockBadgeGradient(topColor string, topOpacity float64, bottomColor string, bottomOpacity float64) style.LinearGradient {
	return stockLinearGradient("texture", topColor, topOpacity, bottomColor, bottomOpacity)
}

func stockLinearGradient(id string, topColor string, topOpacity float64, bottomColor string, bottomOpacity float64) style.LinearGradient {
	return style.LinearGradient{
		ID:     id,
		Top:    style.Point{X: 50, Y: 0},
		Bottom: style.Point{X: 50, Y: 100},
		Stops: []style.Stop{
			style.Stop{
				Color:   bottomColor,
				Opacity: bottomOpacity,
				Offset:  0,
			},
			style.Stop{
				Color:   topColor,
				Opacity: topOpacity,
				Offset:  100,
			},
		},
	}
}

func stockTextFilter() style.Filter {
	return style.Filter{ID: "shadow"}
}

func stockText(text string, id string, origin geo.Coordinate) shapes.Text {
	return shapes.Text{
		ID:    id,
		Value: text,
		Font: shapes.Font{
			Family: "Verdana",
			Size:   10,
			Weight: "normal",
		},
		Origin: origin,
	}
}

func makeTextID(text string, element string) string {
	return fmt.Sprintf("%s-%s-text", strings.Replace(text, " ", "-", 1), element)
}
