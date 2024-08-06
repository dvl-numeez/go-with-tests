package propertybasedtest

import "testing"


func TestRomanNumerals(t *testing.T) {
	got:=ConvertToRoman(1)
	wanted:="I"
	if got!=wanted{
		t.Errorf("Wanted : %q Got : %q",wanted,got)
	}
}