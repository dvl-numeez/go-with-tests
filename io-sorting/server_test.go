package iosorting

import (

	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	
	"testing"
)


func TestFileSystemStore(t *testing.T) {

	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store,err := NewFileSystemStore(database)
		if err!=nil{

		}

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		got = store.GetLeague()
		assertLeague(t, got, want)
		assertLeague(t, got, want)
	})
	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store,err:=NewFileSystemStore(database)
		assertNoError(t,err)
	
		got := store.GetPlayerScore("Chris")
	
		want := 33
	
		assertScoreEqual(t,got,want)
	})
	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
	
		store,err:= NewFileSystemStore(database)
	
		store.RecordWin("Chris")
	
		got := store.GetPlayerScore("Chris")
		want := 34
		assertNoError(t,err)
		assertScoreEqual(t, got, want)
	})
	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
	
		store,err:=NewFileSystemStore(database)
		assertNoError(t,err)
	
		store.RecordWin("Pepper")
	
		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEqual(t, got, want)
	})
	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()
	
		_, err := NewFileSystemStore(database)
	
		assertNoError(t, err)
	})
	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
	
		store, err := NewFileSystemStore(database)
	
		assertNoError(t, err)
	
		got := store.GetLeague()
	
		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}
		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}

func AssertResponseBody(t testing.TB,got,want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	
}

func NewGetScoreRequest(name string)*http.Request{
	url:=fmt.Sprintf("/players/%s",name)
	req,_:= http.NewRequest(http.MethodGet, url, nil)
	return req
}
func NewPostRequest(name string)*http.Request{
	url:=fmt.Sprintf("/players/%s",name)
	req,_:=http.NewRequest(http.MethodPost, url, nil)
	return req
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return
}

func assertLeague(t testing.TB, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

func assertScoreEqual(t testing.TB,got ,want int){
	t.Helper()
	if got!=want{
		t.Errorf("Actual : %d , Expected : %d",got ,want)
	}
}

func createTempFile(t testing.TB,initialDataString string)(*os.File,func()){
	t.Helper()
	tmpFile,err:=os.CreateTemp("","db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}
	tmpFile.Write([]byte(initialDataString))
	removeFile:=func(){
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}
	return tmpFile,removeFile
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
