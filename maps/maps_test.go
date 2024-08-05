package maps

import "testing"

func TestSear(t *testing.T){
	dictionary:= Dictionary{"test":"this is a test"}
	got:=dictionary.Search("test")
	want:="this is a test"
	assertStrings(t,got,want)
	
}

func assertStrings(t testing.TB,got,want string){
	t.Helper()
	if got!=want{
		t.Errorf("Got : %s Wanted %s",got,want)
	}
}