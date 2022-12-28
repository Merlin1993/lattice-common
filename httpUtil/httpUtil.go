package httpUtil

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	Client *http.Client
}

func (f *HttpClient) Get(url string) (string, error) {
	resp, err := f.Client.Get(url)
	if err != nil {
		return "", errors.New(err.Error())
	}
	defer resp.Body.Close()
	// 全读取，会出现joule不足错误
	//result, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return "", errors.New(err.Error())
	//}
	//return string(result), nil

	// 设置读取内容长度，避免joule不足
	buf := make([]byte, 4096)
	result, err := resp.Body.Read(buf)
	if err == nil || err == io.EOF {
		return string(buf[:result]), nil

	}
	return "", errors.New(err.Error())
}

func (f *HttpClient) Fetch(url string, data interface{}) ([]byte, error) {
	jsonStr, _ := json.Marshal(data)
	resp, err := f.Client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 全读取，是否会出现joule不足错误
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil

	// 设置读取内容长度，避免joule不足
	//buf := make([]byte, 4096)
	//result, err := resp.Body.Read(buf)
	//if err != nil {
	//	return nil, err
	//}
	//return buf[:result], nil
}
