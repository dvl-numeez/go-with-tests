package sync

import "testing"


func TestCount(t *testing.T){
	t.Run("Incrementing the counter 3 times and leave it with that value",func(t *testing.T){
		counter:= Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCount(t,&counter,3)
	})
}



func assertCount( t *testing.T,counter *Counter,expectedValue int ){
	t.Helper()
	if counter.value!=expectedValue{
		t.Errorf("Got %d wanted %d",counter.Value(),expectedValue)
	}
}

