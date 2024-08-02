package methods

import "testing"

func TestPermieter(t *testing.T){
	rectangle:= Rectangle{
		Width: 10.0,
		Height: 10.0,
	}
	actual:=rectangle.Perimeter()
	wanted:= 40.0
	if actual!=wanted{
		t.Errorf("Actual %.2f Expected %.2f",actual,wanted)
	}

}


func TestArea(t *testing.T){
	t.Run("Rectangles",func(t *testing.T){
		rectangle:= Rectangle{
			Width: 12.0,
			Height: 6.0,
		}
		actual:=rectangle.Area()
		wanted:=72.0
		if actual!=wanted{
			t.Errorf("Actual %.2f Expected %.2f",actual,wanted)
		}
	})
}