package mocking

import (
	"bytes"
	
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {
	t.Run("Testing the countdown function",func(t *testing.T){
		buffer:=&bytes.Buffer{}
		spySleeper:=&SpySleeper{}
		Countdown(buffer,spySleeper)
		got:=buffer.String()
		wanted:=`3
2
1
Go!`
		if got!=wanted{
			t.Errorf("Got : %s Wanted : %s",got,wanted)
		}
		if spySleeper.Calls!=3{
			t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
		}
	})
	
	
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}