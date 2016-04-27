package client

import (
	"fmt"
	"testing"
)

func Test_CreateToken_1(t *testing.T) {
	token := CreateToken("groups", "admin", "admin")
	fmt.Println(token)
	t.Log(token) //记录一些你期望记录的信息
}

func Test_UrlMethodMap(t *testing.T) {
	urls := []string{
		"/api/user",
		"/api/log",
		"/api/user/test/123",
	}
	for _, j := range urls {
		method, err := UrlMethodMap(j)
		if err != nil {
			t.Fatal(err)
		} else {
			t.Log(method)
		}
	}
}
