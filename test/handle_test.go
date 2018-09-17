package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bouk/monkey"

	"test_inline/handle"
)

func TestHandle(t *testing.T)  {
	monkey.Patch(handle.Max, handle.MaxAddOne)

	result := handle.Max(1,2)
	if result == 2 {
		t.Error(result)
	} else if result == 3 {
		t.Log(result)
	}
}

func TestHttp(t *testing.T) {

	monkey.Patch(handle.Max, handle.MaxAddOne)
	server := httptest.NewServer(http.HandlerFunc(handle.SayMax))

	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	//fmt.Println(string(body))
	result := string(body)
	if result == "2" {
		t.Error(result)
	} else if result == "3" {
		t.Log(result)
	}
}