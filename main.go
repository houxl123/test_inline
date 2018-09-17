package main

import (
	"net/http"
	"test_inline/handle"
)


func main() {
	http.HandleFunc("/", handle.SayMax)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}