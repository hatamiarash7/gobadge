package style

import (
	"fmt"
	"strings"

	"github.com/hatamiarash7/gobadge/svg/shapes"
)

// Text is the text style for badge's text
type Text struct {
	ID     string
	Filter Filter
	Text   shapes.Text
}

// Vectorize the text
func (text Text) Vectorize() string {
	uses := strings.Join([]string{
		fmt.Sprintf(`<use fill-opacity="1" filter="url(#%s)" xlink:href="#%s"></use>`, text.Filter.ID, text.Text.ID),
		fmt.Sprintf(`<use fill-rule="evenodd" xlink:href="#%s"></use>`, text.Text.ID),
		fmt.Sprintf(`<use fill-opacity="1" xlink:href="#%s"></use>`, text.Text.ID),
	}, "\n")

	template := `<g id="%s" fill="#FFFFFF">%s</g>`

	return fmt.Sprintf(template, text.ID, uses)
}
