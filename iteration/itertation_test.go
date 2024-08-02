package iteration

import "testing"


func TestIteration(t *testing.T){
	actual:=Repeat("a")
	expected:="aaaaa"
	if actual!=expected{
		t.Errorf("Actual : %q Expected:%q",actual,expected)
	}
}