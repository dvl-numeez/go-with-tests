package select_statement

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)


func TestRacer(t *testing.T) {
	t.Run("comparing the speed of fast and slow url ",func(t *testing.T) {
		slowServer:=makeDelayedServer(20*time.Millisecond)
	fastServer:=makeDelayedServer(0*time.Millisecond)
	
	slowUrl:=slowServer.URL
	fastUrl:=fastServer.URL
	want := fastUrl
	got,_ := Racer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	})
	t.Run("returns an error if it takes more than 10 seconds",func(t *testing.T){
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
	

}


func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}