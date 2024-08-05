package reflection

import "testing"


func TestWalk(t *testing.T) {
	t.Run("Testing walk function",func(t *testing.T){
		expected:="Numeez"
		var got []string
		x:=struct{
			Name string
		}{expected}
		walk(x,func(input string){
			got = append(got, input)
		})
		if len(got) != 1 {
			t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
		}
	})
}