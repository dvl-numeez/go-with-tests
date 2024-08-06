package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)


type SpyStore struct{
	response string
	cancelled bool
	t *testing.T
}
func (s *SpyStore)assertWasCancelled(){
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("store was not told to cancel")
	}
}
func (s *SpyStore)assertWasNotCancelled(){
	s.t.Helper()
	if s.cancelled {
		s.t.Error("store was told to cancel")
	}
}
func (s *SpyStore)Fetch()string{
	time.Sleep(10*time.Millisecond)
	return s.response
}
func (s *SpyStore)Cancel(){
	s.cancelled=true
}


func TestServer(t *testing.T) {
	t.Run("Testing happy path",func(t *testing.T){
	data:="Hello World"
	store:=&SpyStore{response:data}
	srv:=Server(store)
	req:=httptest.NewRequest(http.MethodGet,"/",nil)
	res:=httptest.NewRecorder()
	srv.ServeHTTP(res,req)
	if res.Body.String()!=data{
		t.Errorf("got : %q expected : %q ",res.Body.String(),data)
	}
	store.assertWasNotCancelled()
	})

	t.Run("tells store to cancel the work when the fetch is not completed",func(t *testing.T){
		data:="Hello"
		store:=&SpyStore{response:data}
		srv:=Server(store)
		request:=httptest.NewRequest(http.MethodGet,"/",nil)
		cancellingCtx,cancel:=context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond,cancel)
		request  = request.WithContext(cancellingCtx)
		response:=httptest.NewRecorder()
		srv.ServeHTTP(response,request)
		store.assertWasCancelled()
	})
	

}