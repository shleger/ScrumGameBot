package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/shleger/ScrumGameBot/datastore"
	"github.com/shleger/ScrumGameBot/domain"
)

var (
	server  *httptest.Server
	reader  io.Reader
	userURL string

	jsonPost       = `{"update_id":501758830,"message":{"message_id":10,"from":{"id":389814768,"first_name":"Andrew","last_name":"Sch.","language_code":"en-US"},"chat":{"id":389814768,"first_name":"Andrew","last_name":"Sch.","type":"private"},"date":1503348017,"text":"Aaa123"}}`
	jsonPostNorris = `{"name":"Norris"}`

	testTask    *domain.Task
	testTaskKey string
)

func init() {
	server = httptest.NewServer(http.HandlerFunc(Hook))
	//	defer server.Close()
	userURL = fmt.Sprintf("%s/hook", server.URL)
}

func TestMapping(t *testing.T) {

	defer server.Close()

	reader = strings.NewReader(jsonPost)
	req, err := http.NewRequest("POST", userURL, reader)

	if err != nil {
		t.Errorf("Something happened %v", err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Resp Error code %v", resp.StatusCode)
	}

}

type PropServiceTest struct {
	DbSrv datastore.PropsService
}

func (c PropServiceTest) EchoTask(key string) string {
	return "TestStub+" + key
}

func (c PropServiceTest) GetKey(string) string {
	return "TestStub1"
}

func (c PropServiceTest) GetToken(string, string) string {
	return "TestTokenInAuthNameSpace"
}

func (c PropServiceTest) PutKey(key string, task *domain.Task) {
	testTask = task
	testTaskKey = key

}

func check(expected string, actual string, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestPutTask(t *testing.T) {
	app := &App{DbSrv: &PropServiceTest{}}

	req, _ := http.NewRequest("GET", "/taskput", nil)
	w := httptest.NewRecorder()

	app.PutTask(w, req)

	actual := w.Body.String()
	expected := "hello from db"

	check(expected, actual, t)
	check("sampletask3", testTaskKey, t)
	check("Buy milk3", testTask.Description, t)

}

func TestGetTask(t *testing.T) {
	app := &App{DbSrv: &PropServiceTest{}}

	req, _ := http.NewRequest("GET", "/taskget", nil)
	w := httptest.NewRecorder()

	app.GetTask(w, req)

	actual := w.Body.String()
	expected := "TestTokenInAuthNameSpace"

	check(expected, actual, t)

}

func TestEcho(t *testing.T) {

	app := &App{DbSrv: &PropServiceTest{}}

	req, _ := http.NewRequest("GET", "/echo", nil)
	w := httptest.NewRecorder()

	app.Echo(w, req)

	actual := w.Body.String()
	expected := "TestStub+echoKey"

	check(expected, actual, t)

}
