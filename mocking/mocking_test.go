package mocking

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	buffer:=&bytes.Buffer{}
	CountDown(buffer)
	got:=buffer.String()
	wanted:="3"
	if got!=wanted{
		t.Errorf("Got : %s Wanted : %s",got,wanted)
	}
}