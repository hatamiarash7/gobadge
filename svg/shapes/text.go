package shapes

import (
	"fmt"
	"strings"

	geo "github.com/hatamiarash7/gobadge/svg/geometry"
)

// Text is the text shape for badge's text
type Text struct {
	ID     string
	Value  string
	Font   Font
	Origin geo.Coordinate
}

// Font of our text
type Font struct {
	Family string
	Size   int
	Weight string
}

// Vectorize the text
func (text Text) Vectorize() string {
	elements := strings.Join([]string{
		fmt.Sprintf(`id="%s"`, text.ID),
		fmt.Sprintf(`font-family="%s"`, text.Font.Family),
		fmt.Sprintf(`font-size="%d"`, text.Font.Size),
		fmt.Sprintf(`font-weight="%s"`, text.Font.Weight),
	}, " ")

	contents := fmt.Sprintf(`<tspan %s>%s</tspan>`, text.Origin.Vectorize(), text.Value)
	template := `<text %s>%s</text>`

	return fmt.Sprintf(template, elements, contents)
}
