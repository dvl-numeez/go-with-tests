package methods

type Rectangle struct {
	Width float64
	Height float64
}
func(rectangle Rectangle) Perimeter()float64{
	return 2*(rectangle.Width+rectangle.Height)
}