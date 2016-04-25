package client

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"time"
)

type Request interface {
	SetHeader()
	SetAuth()
	DoRequest()
}

type ClientOpts struct {
	Url       string
	AccessKey string
	SecretKey string
	Timeout   time.Duration
}

type BaseClient struct {
	Opts *ClientOpts
}

type Op struct {
	Name string
}

const (
	DefaultTimeOut time.Duration = time.Second * 10
)

var (
	Get    = Op{Name: "GET"}
	Head   = Op{Name: "HEAD"}
	Delete = Op{Name: "DELETE"}
	Put    = Op{Name: "PUT"}
	Post   = Op{Name: "POST"}
	Update = Op{Name: "UPDATE"}
)

func NewRequest(url, op string, timeout time.Duration, body io.Reader) (*http.Request, error) {
	if len(url) == 0 || len(op) == 0 {
		return &http.Request{}, errors.New("invalid argument")
	}
	req, err := http.NewRequest(op, url, body)
	if err != nil {
		return &http.Request{}, err
	}
	return req, nil
}

func (c BaseClient) DoAction(path string, op Op) (resp *http.Response, err error) {
	Timeout := c.Opts.Timeout
	if Timeout == 0 {
		Timeout = DefaultTimeOut
	}
	client := &http.Client{Timeout: Timeout}

	req, err := http.NewRequest(op.Name, c.Opts.Url+path, nil)
	if err != nil {
		panic(err.Error())
	}
	//参考http://stackoverflow.com/questions/17714494/golang-http-request-results-in-eof-errors-when-making-multiple-requests-successi
	//EOF错误
	req.Close = true

	req.SetBasicAuth(c.Opts.AccessKey, c.Opts.SecretKey)
	resp, err = client.Do(req)
	return

}

func (c BaseClient) DoPost(path string, data []byte) (resp *http.Response, err error) {
	body := bytes.NewReader(data)

	Timeout := c.Opts.Timeout
	if Timeout == 0 {
		Timeout = DefaultTimeOut
	}
	client := &http.Client{Timeout: Timeout}

	req, err := http.NewRequest("POST", c.Opts.Url+path, body)
	if err != nil {
		panic(err.Error())
	}
	req.Close = true

	req.SetBasicAuth(c.Opts.AccessKey, c.Opts.SecretKey)
	resp, err = client.Do(req)
	return

}
