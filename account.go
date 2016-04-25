package client

import (
	"bytes"
	"net/http"
)

func (c *ApphouseClient) AddAccount(byteContent []byte) (interface{}, error) {
	url := "/api/account/add"
	method, err := UrlToMethod(url)
	if err != nil {
		return "", err
	}

	fullUrl := "http://" + c.ServerIp + ":" + c.ServerPort + url
	token := CreateToken(method, c.AccessKey, c.SecretKey)
	req, err := http.NewRequest("POST", fullUrl, bytes.NewReader(byteContent))
	req.Header.Add("token", token)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	retContent, err := parseResponse(resp)
	if err != nil {
		return "", err
	}
	return retContent, nil
}

func (c *ApphouseClient) DelAccount(byteContent []byte) (interface{}, error) {
	url := "/api/account/delete"
	method, err := UrlToMethod(url)

	if err != nil {
		return &http.Response{}, err
	}

	fullUrl := "http://" + c.ServerIp + ":" + c.ServerPort + url
	token := CreateToken(method, c.AccessKey, c.SecretKey)
	req, err := http.NewRequest("POST", fullUrl, bytes.NewReader(byteContent))
	req.Header.Add("token", token)

	client := &http.Client{}
	resp, err := client.Do(req)

	retContent, err := parseResponse(resp)
	if err != nil {
		return "", err
	}
	return retContent, nil

}

func (c *ApphouseClient) GetAccountInfo(byteContent []byte) (interface{}, error) {
	url := "/api/account"
	method, err := UrlToMethod(url)
	if err != nil {
		return &http.Response{}, err
	}

	userId := string(byteContent)
	fullUrl := "http://" + c.ServerIp + ":" + c.ServerPort + url + "/" + userId
	token := CreateToken(method, c.AccessKey, c.SecretKey)
	req, err := http.NewRequest("GET", fullUrl, nil)
	req.Header.Add("token", token)

	client := &http.Client{}
	resp, err := client.Do(req)

	retContent, err := parseResponse(resp)
	if err != nil {
		return "", err
	}
	return retContent, nil
}

func (c *ApphouseClient) GetAccounts() (interface{}, error) {
	url := "/api/accounts"
	method, err := UrlToMethod(url)
	if err != nil {
		return &http.Response{}, err
	}

	fullUrl := "http://" + c.ServerIp + ":" + c.ServerPort + url
	token := CreateToken(method, c.AccessKey, c.SecretKey)
	req, err := http.NewRequest("GET", fullUrl, nil)
	req.Header.Add("token", token)

	client := &http.Client{}
	resp, err := client.Do(req)

	retContent, err := parseResponse(resp)
	if err != nil {
		return "", err
	}
	return retContent, nil
}
