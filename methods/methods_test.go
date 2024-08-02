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