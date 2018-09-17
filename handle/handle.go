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