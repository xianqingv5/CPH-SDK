package httphelper2

import (
	"CPH-SDK/conf"
	"CPH-SDK/core"
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(uri string) ([]byte, error) {
	s := core.Signer{
		Key:    conf.Config.Huawei.Key,
		Secret: conf.Config.Huawei.Secret,
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("X-Auth-Token", authtoken.Authtoken())
	s.Sign(req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func HttpPost(uri string, data []byte) ([]byte, error) {
	s := core.Signer{
		Key:    conf.Config.Huawei.Key,
		Secret: conf.Config.Huawei.Secret,
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", uri, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	s.Sign(req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func HttpPut(uri string, data []byte) ([]byte, error) {
	s := core.Signer{
		Key:    conf.Config.Huawei.Key,
		Secret: conf.Config.Huawei.Secret,
	}
	client := &http.Client{}
	req, err := http.NewRequest("PUT", uri, strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	s.Sign(req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
