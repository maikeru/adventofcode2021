package point

import "fmt"

type Point struct {
	X int
	Y int
}

func (point *Point) String() string {
	return fmt.Sprintf("Point X: %d, Y: %d", point.X, point.Y)
}
