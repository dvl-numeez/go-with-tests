package hello

import "testing"


func TestHello(t *testing.T){
	actual:=Hello("Numeez","")
	expected:="Hello Numeez"
	assertCorrectMessage(t,actual,expected)
	
	t.Run("Say hello world when the empty string is passed",func(t *testing.T){
		actual:=Hello("","")
		expected:="Hello World"
		assertCorrectMessage(t,actual,expected)
	})
	t.Run("Hello in Spanish",func(t *testing.T) {
		actual:=Hello("Numeez","Spanish")
		expected:="Hola Numeez"
		assertCorrectMessage(t,actual,expected)
	})
	t.Run("Hello in French",func(t *testing.T) {
		actual:=Hello("Numeez","French")
		expected:="Bonjour Numeez"
		assertCorrectMessage(t,actual,expected)
	})

}

func assertCorrectMessage(t testing.TB,actual,expected string){
	t.Helper()
	if actual!=expected{
		t.Errorf("Actual : %q Expected : %q",actual,expected)
	}
}