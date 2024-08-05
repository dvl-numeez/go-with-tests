package mocking

import (
	"fmt"
	"io"
	"time"

)


const finalWord = "Go!"
const countdownStart = 3
const write = "write"
const sleep = "sleep"

type SpyCountDownOperation struct {
	Calls []string
}

func  (s SpyCountDownOperation)Sleep(){
	s.Calls = append(s.Calls, sleep)
}
func  (s SpyCountDownOperation)Write(p []byte)(n int ,err error){
	s.Calls = append(s.Calls, write)
	return
}
type Sleeper interface{
	Sleep()
}
type SpySleeper struct{
	Calls int
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
type SpyTime struct{
	durationSlept time.Duration
}
func (s *SpyTime)Sleep(duration time.Duration){
	s.durationSlept = duration
}
func(s *SpySleeper)Sleep(){
	s.Calls++
}

func Countdown(out io.Writer,sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		
		
	}
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		
		
	}
	fmt.Fprint(out, finalWord)
}