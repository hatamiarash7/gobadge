package shapes

import (
	"fmt"
)

// Point is a XY coordinate
type Point Coordinate

// Vectorize our point
func (point Point) Vectorize() string {
	x := fmt.Sprintf(`x="%f"`, point.X)
	y := fmt.Sprintf(`y="%f"`, point.Y)

	return fmt.Sprintf("%s %s", x, y)
}
