package client

import (
	"crypto/hmac"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

var (
	//用于Url和method方法名的映射
	UrlMethodMap = make(map[string]string)
)

type ApphouseClient struct {
	BaseClient
	Account AccountOperations
}

func setupBaseClient(apphouseClient *BaseClient, opts ClientOpts) error {
	apphouseClient.Opts.BaseUrl = opts.BaseUrl
	apphouseClient.Opts.AccessKey = opts.AccessKey
	apphouseClient.Opts.SecretKey = opts.SecretKey
	return nil
}

func NewApphouseClient(opts ClientOpts) (*ApphouseClient, error) {

	c := &ApphouseClient{}
	c.Account = newAccountClient(c)
	setupBaseClient(&c.BaseClient, opts)
	return c, nil
}

func parseResponse(resp *http.Response) (interface{}, error) {

	defer func() {
		if resp != nil || resp.Body != nil {
			resp.Body.Close()
		}
	}()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	var result Response
	err = json.Unmarshal(content, &result)
	if err != nil {
		return []byte(""), err
	}

	if result.Result != 0 {
		return []byte(""), errors.New(result.Message)
	}

	return result.Content, nil
}

//method:方法名称， accessKey :密码 uuid :用户名
// /api/tag/delete -- method  delete_tag
func CreateToken(method string, secretKey string, accessKey string) string {

	timeStamp := time.Now().Unix() * 1000
	msg := fmt.Sprintf("<%d><%s>", timeStamp, method)

	token := make(map[string]string)
	token["timestamp"] = fmt.Sprintf("%d", timeStamp)
	hmacHash := hmac.New(sha1.New, []byte(secretKey))
	hmacHash.Write([]byte(msg))
	expectedMac := hmacHash.Sum(nil)
	hexDigest := hex.EncodeToString(expectedMac)

	//encoder := base64.NewEncoding(hmac.New(s, accessKey))
	encoded := base64.URLEncoding.EncodeToString([]byte(hexDigest))
	token["hash"] = encoded

	token["uuid"] = base64.StdEncoding.EncodeToString([]byte(accessKey))

	return token["uuid"] + "%%" + token["timestamp"] + "%%" + token["hash"]
}

func UrlToMethod(url string) (string, error) {
	/*
		if len(url) == 0 {
			return "", errors.New("url is empty")
		}
		urlTrim := strings.TrimPrefix(url, "/api/")
		var method string
		slice := strings.SplitN(urlTrim, "/", -1)
		i := len(slice) - 1
		for {
			if i < 0 {
				break
			}
			if i != (len(slice) - 1) {
				method = method + "_" + slice[i]
			} else {
				method = slice[i]
			}
			i = i - 1
		}
		return method, nil
	*/
	method, ok := UrlMethodMap[url]
	if !ok {
		return "", errors.New("method doesn't exists for " + url)
	}
	return method, nil

}

func init() {
	UrlMethodMap["/api/account/add"] = "add_account"
	UrlMethodMap["/api/account/delete"] = "delete_account"
	UrlMethodMap["/api/accounts"] = "accounts"
	UrlMethodMap["/api/account"] = "account"
	UrlMethodMap["/api/account/update"] = "update_account"

}
