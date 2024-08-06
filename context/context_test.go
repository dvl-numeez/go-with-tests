package context

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


type SpyStore struct{
	response string
}

func (s *SpyStore)Fetch()string{
	return s.response
}


func TestServer(t *testing.T) {
	data:="Hello World"
	srv:=Server(&SpyStore{data})
	req:=httptest.NewRequest(http.MethodGet,"/",nil)
	res:=httptest.NewRecorder()
	srv.ServeHTTP(res,req)
	if res.Body.String()!=data{
		t.Errorf("got : %q expected : %q ",res.Body.String(),)
	}

}