package cph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"authtoken"
)

func ListCloudPhoneServers() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/phones?offset=0&limit=10", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", authtoken.Authtoken())
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(body))
}
