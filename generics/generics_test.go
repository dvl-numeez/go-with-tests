package generics

import "testing"



func TestAssertFunctions(t *testing.T){
	t.Run("Asserting functions",func(t *testing.T){
		AssertEqual(t,1,1)
		AssertNotEqual(t,1,2)
	})
	t.Run("assert on strings",func(t *testing.T){
		AssertEqual(t, "hello", "hello")
	AssertNotEqual(t, "hello", "Numeez")
	})
}

func AssertEqual[T comparable](t *testing.T,got,want T){
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}


func AssertNotEqual[T comparable](t *testing.T,got,want T){
	t.Helper()
	if got == want {
		t.Errorf("didn't want %d", got)
	}
}