package client

import (
	"encoding/json"
	"errors"
)

type AccountClient struct {
	apphouseClient *ApphouseClient
}

type Account struct {
	Id       string `json:"_id"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Desc     string `json:"desc"`
	Source   string `json:"source"`
}

type AccountWithoutId struct {
	Password string `json:"password"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Desc     string `json:"desc"`
	Source   string `json:"source"`
}

type Uid struct {
	Id string `json:"_id"`
}

type AccountCollection struct {
	Data []Account `json:"account"`
}

type AccountOperations interface {
	Create(info interface{}) (*Account, error)
	Delete(container *Account) error
	Info(tag interface{}) (*Account, error)
	List(groupId string) (*AccountCollection, error)
	Update(existing *Account, updates interface{}) error
}

func newAccountClient(apphouseClient *ApphouseClient) *AccountClient {
	return &AccountClient{
		apphouseClient: apphouseClient,
	}
}

func (c *AccountClient) Create(info interface{}) error {
	url := "/api/account/add"
	method, err := UrlToMethod(url)
	if err != nil {
		return err
	}

	var fullUrl string
	fullUrl = c.apphouseClient.Opts.BaseUrl + url

	token := CreateToken(method, c.apphouseClient.Opts.AccessKey, c.apphouseClient.Opts.SecretKey)

	headers := make(map[string]string)
	headers["token"] = token

	var resp Response
	err = c.apphouseClient.Post(fullUrl, headers, info, &resp)
	if err != nil {
		return err
	}

	if resp.Result != 0 {
		return errors.New(resp.Message)
	}
	return nil
}

func (c *AccountClient) Delete(container *Account) error {
	url := "/api/account/delete"
	method, err := UrlToMethod(url)
	if err != nil {
		return err
	}

	var fullUrl string
	fullUrl = c.apphouseClient.Opts.BaseUrl + url

	token := CreateToken(method, c.apphouseClient.Opts.AccessKey, c.apphouseClient.Opts.SecretKey)

	headers := make(map[string]string)
	headers["token"] = token

	uid := new(Uid)
	uid.Id = container.Id
	var resp Response
	err = c.apphouseClient.Post(fullUrl, headers, uid, &resp)
	if err != nil {
		return err
	}

	if resp.Result != 0 {
		return errors.New(resp.Message)
	}

	return nil
}

func (c *AccountClient) Info(info interface{}) (*Account, error) {
	url := "/api/account"
	method, err := UrlToMethod(url)
	if err != nil {
		return &Account{}, err
	}

	uid, ok := info.(Uid)
	if !ok {
		return &Account{}, errors.New("invalid type param")
	}

	var fullUrl string
	fullUrl = c.apphouseClient.Opts.BaseUrl + url + "/" + uid.Id
	token := CreateToken(method, c.apphouseClient.Opts.AccessKey, c.apphouseClient.Opts.SecretKey)

	headers := make(map[string]string)
	headers["token"] = token

	var resp Response
	err = c.apphouseClient.Get(fullUrl, headers, &resp)
	if err != nil {
		return &Account{}, err
	}

	if resp.Result != 0 {
		return &Account{}, errors.New(resp.Message)
	}

	var account Account
	err = json.Unmarshal(resp.Content, &account)
	if err != nil {
		return &Account{}, err
	}

	return &account, nil
}

func (c *AccountClient) List(groupsId string) (*AccountCollection, error) {
	url := "/api/accounts"
	method, err := UrlToMethod(url)
	if err != nil {
		return &AccountCollection{}, err
	}

	var fullUrl string
	if len(groupsId) != 0 {
		fullUrl = c.apphouseClient.Opts.BaseUrl + url + "/" + groupsId
	} else {
		fullUrl = c.apphouseClient.Opts.BaseUrl + url
	}
	token := CreateToken(method, c.apphouseClient.Opts.AccessKey, c.apphouseClient.Opts.SecretKey)

	headers := make(map[string]string)
	headers["token"] = token
	var resp Response
	err = c.apphouseClient.Get(fullUrl, headers, &resp)
	if err != nil {
		return &AccountCollection{}, err
	}

	if resp.Result != 0 {
		return &AccountCollection{}, errors.New(resp.Message)
	}

	var col AccountCollection
	err = json.Unmarshal(resp.Content, &col)
	if err != nil {
		return &AccountCollection{}, err
	}

	return &col, nil
}

func (c *AccountClient) Update(existing *Account, updates interface{}) error {

	url := "/api/account/update"
	method, err := UrlToMethod(url)
	if err != nil {
		return err
	}

	var fullUrl string
	fullUrl = c.apphouseClient.Opts.BaseUrl + url + "/" + existing.Id
	token := CreateToken(method, c.apphouseClient.Opts.AccessKey, c.apphouseClient.Opts.SecretKey)

	headers := make(map[string]string)
	headers["token"] = token

	var resp Response
	err = c.apphouseClient.Post(fullUrl, headers, Update, &resp)
	if err != nil {
		return err
	}

	if resp.Result != 0 {
		return errors.New(resp.Message)
	}

	return nil
}
