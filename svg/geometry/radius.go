package geometry

import (
	"fmt"
)

// Radius is the radius of our shape
type Radius Coordinate

// Vectorize the radius
func (radius Radius) Vectorize() string {
	x := fmt.Sprintf(`rx="%f"`, radius.X)
	y := fmt.Sprintf(`ry="%f"`, radius.Y)

	return fmt.Sprintf("%s %s", x, y)
}
