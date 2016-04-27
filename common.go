package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type ClientOpts struct {
	BaseUrl   string
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

func (c *BaseClient) newHttpClient() *http.Client {
	return &http.Client{}
}

func (c *BaseClient) newHttpRequset() *http.Request {
	return &http.Request{}
}

func (c *BaseClient) setHeader(r *http.Request, headers map[string]string) {
	for k, v := range headers {
		r.Header.Set(k, v)
	}
}

func (c *BaseClient) doModify(method string, url string, headers map[string]string, createObject interface{}, respObject interface{}) error {
	byteContent, err := json.Marshal(createObject)
	if err != nil {
		return err
	}

	client := c.newHttpClient()
	req, err := http.NewRequest(method, url, bytes.NewBuffer(byteContent))
	if err != nil {
		return nil
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		msg := fmt.Sprintf("StatusCode:%v,Status:%v", resp.StatusCode, resp.Status)
		return errors.New(msg)
	}
	byteContent, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(byteContent) > 0 {
		return json.Unmarshal(byteContent, respObject)
	}
	return nil
}

func (c *BaseClient) Post(url string, headers map[string]string, createObj interface{}, respObject interface{}) error {
	return c.doModify("POST", url, headers, createObj, respObject)
}
func (c *BaseClient) Delete(url string, headers map[string]string, createObj interface{}, respObject interface{}) error {
	return c.doModify("DELETE", url, headers, createObj, respObject)
}

func (c *BaseClient) doGet(method string, url string, headers map[string]string, respObject interface{}) error {

	client := c.newHttpClient()
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		msg := fmt.Sprintf("StatusCode:%v,Status:%v", resp.StatusCode, resp.Status)
		return errors.New(msg)
	}
	byteContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(byteContent) > 0 {
		return json.Unmarshal(byteContent, respObject)
	}
	return nil
}

func (c *BaseClient) Get(url string, headers map[string]string, respObject interface{}) error {
	return c.doGet("GET", url, headers, respObject)
}

/*
func (c *BaseClient) doDelete(url string)error {
	 client := c.newHttpClient()
	 req, err :=

}*/

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

/*
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
*/
