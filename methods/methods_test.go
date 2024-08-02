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
	areaTests:=[]struct{
		shape Shape
		want float64
	}{
		{Rectangle{12,6},72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}
	for _,tests:= range areaTests{
		got:=tests.shape.Area()
		expected:=tests.want
		if got!=expected{
			t.Errorf("Got : %.2f Wanted : %.2f",got,expected)
		}
	}
}