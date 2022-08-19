package geometry

import (
	"fmt"
)

// Coordinate is a XY coordinate
type Coordinate struct {
	X float64
	Y float64
}

// Vectorize our coordinate
func (coordinate Coordinate) Vectorize() string {
	x := fmt.Sprintf(`x="%f"`, coordinate.X)
	y := fmt.Sprintf(`y="%f"`, coordinate.Y)

	return fmt.Sprintf("%s %s", x, y)
}
