package shapes

import (
	"fmt"
)

// Path is the SVG path element
type Path struct {
	ID       string
	Fill     string
	FillRule string
	D        string
}

// Vectorize the path
func (path Path) Vectorize() string {
	template := `<path id="%s" fill="%s" fill-rule="%s" d="%s"> </path>`

	return fmt.Sprintf(template, path.ID, path.Fill, path.FillRule, path.D)
}
