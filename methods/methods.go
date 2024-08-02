package methods
import "math"

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct{
	Radius float64
}
type Triangle struct{
	Base float64
	Height float64
}

type Shape interface {
	Area()float64
}


func(rectangle Rectangle) Perimeter()float64{
	return 2*(rectangle.Width+rectangle.Height)
}
func(rectangle Rectangle) Area()float64{
	return rectangle.Width*rectangle.Height
}
func (triangle Triangle)Area()float64{
	return (triangle.Base * triangle.Height) * 0.5
}

func (c Circle)Area()float64{
	return math.Pi * c.Radius * c.Radius
}