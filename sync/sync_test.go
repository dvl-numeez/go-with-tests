package sync

import (
	"sync"
	"testing"
)


func TestCount(t *testing.T){
	t.Run("Incrementing the counter 3 times and leave it with that value",func(t *testing.T){
		counter:= Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCount(t,&counter,3)
	})
	t.Run("Check whether it concurrency safe",func(t *testing.T){
		wantedCount:=1000
		counter:=Counter{}
		var wg sync.WaitGroup
		wg.Add(wantedCount)
		for i:=0;i<1000;i++{
			go func ()  {
				counter.Inc()	
				wg.Done()
			}()
		}
		wg.Wait()

		assertCount(t,&counter,wantedCount)
	})
}



func assertCount( t *testing.T,counter *Counter,expectedValue int ){
	t.Helper()
	if counter.value!=expectedValue{
		t.Errorf("Got %d wanted %d",counter.Value(),expectedValue)
	}
}

