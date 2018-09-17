# test_inline

This is a demo for testing inline functions, go1.11 version.
Execute test cases with `go test -gcflags=all=-l test_inline/test`
when `fmt.println` in `handle.Max` opens Both test cases succeeded.
```
func Max(m,n int) int {
	fmt.Println("Max")
	if m > n {
		return m
	}
	return n
}
```

When closed, `TestHandle` passed, `TestHttp` failed. 
```
func Max(m,n int) int {
	//fmt.Println("Max")
	if m > n {
		return m
	}
	return n
}
```
This happens because the inline is not disabled in the http service, so `-gcflags=all=-l` Is there a bug in go1.11?

#### handle.go
```
package handle

import (
	"fmt"
	"net/http"
)

func Max(m,n int) int {
	//fmt.Println("Max")
	if m > n {
		return m
	}
	return n
}

func MaxAddOne(m, n int) int {
	if m > n {
		return m + 1
	}
	return n + 1
}

func SayMax(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d", Max(1,2))
}
```
#### handle_test.go
```
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
```
