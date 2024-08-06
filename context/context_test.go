package context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyResponseWriter struct{
	written bool
}

func (sr *SpyResponseWriter)Header() http.Header{
	sr.written=true
	return nil
}
func (sr *SpyResponseWriter)Write([]byte) (int, error){
	sr.written=true
	return 0,errors.New("not implemented")
}
func (sr *SpyResponseWriter)WriteHeader(statusCode int){
	sr.written=true
}


type SpyStore struct{
	response string
	t *testing.T
}
// func (s *SpyStore)assertWasCancelled(){
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Error("store was not told to cancel")
// 	}
// }
// func (s *SpyStore)assertWasNotCancelled(){
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Error("store was told to cancel")
// 	}
// }
func (s *SpyStore)Fetch(ctx context.Context)(string,error){
	data:=make(chan string,1)
	go func(){
		var result string
		for _,c:=range s.response{
			select{
			case<-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10*time.Millisecond)
				result+=string(c)
			}
		}
		data<-result
	}()
	select{
	case<-ctx.Done():
		return "",ctx.Err()
	case res:=<-data:
		return res,nil
	}
}
func (s *SpyStore)Cancel(){
	
}


func TestServer(t *testing.T) {
	t.Run("Testing happy path",func(t *testing.T){
	data:="Hello World"
	store:=&SpyStore{response:data,t:t}
	srv:=Server(store)
	req:=httptest.NewRequest(http.MethodGet,"/",nil)
	res:=httptest.NewRecorder()
	srv.ServeHTTP(res,req)
	if res.Body.String()!=data{
		t.Errorf("got : %q expected : %q ",res.Body.String(),data)
	}
	})

	t.Run("tells store to cancel the work when the fetch is not completed",func(t *testing.T){
		data:="Hello"
		store:=&SpyStore{response:data}
		srv:=Server(store)
		request:=httptest.NewRequest(http.MethodGet,"/",nil)
		cancellingCtx,cancel:=context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond,cancel)
		request  = request.WithContext(cancellingCtx)
		response:=&SpyResponseWriter{}
		srv.ServeHTTP(response,request)
		if response.written{
			t.Error("response should not have been written")
		}
		
	})
	

}