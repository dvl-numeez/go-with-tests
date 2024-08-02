package hello

import "testing"


func TestHello(t *testing.T){
	actual:=Hello("Numeez")
	expected:="Hello Numeez"
	if actual!=expected{
		t.Errorf("Actual : %q Expected : %q",actual,expected)
	}
}