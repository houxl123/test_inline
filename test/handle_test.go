package test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/bouk/monkey"

	"test_inline/handle"
)

func TestHandle(t *testing.T) {

	monkey.Patch(handle.Max, handle.MaxAddOne)
	resp, err := http.Get("http://localhost:9090/")
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