package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server  *httptest.Server
	reader  io.Reader
	userURL string

	jsonPost       = `{"update_id":501758830,"message":{"message_id":10,"from":{"id":389814768,"first_name":"Andrew","last_name":"Sch.","language_code":"en-US"},"chat":{"id":389814768,"first_name":"Andrew","last_name":"Sch.","type":"private"},"date":1503348017,"text":"Aaa123"}}`
	jsonPostNorris = `{"name":"Norris"}`
)

func init() {
	server = httptest.NewServer(http.HandlerFunc(Hook))
	//	defer server.Close()
	userURL = fmt.Sprintf("%s/hook", server.URL)
}

func TestHook(t *testing.T) {

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
