package methods
import "math"

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct{
	Radius float64
}

func(rectangle Rectangle) Perimeter()float64{
	return 2*(rectangle.Width+rectangle.Height)
}
func(rectangle Rectangle) Area()float64{
	return rectangle.Width*rectangle.Height
}

func (c Circle)Area()float64{
	return math.Pi * c.Radius * c.Radius
}