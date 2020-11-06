package cphservers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"authtoken"
)

func ListCloudPhoneServers(w http.ResponseWriter, r *http.Request) {
	f := func(offerset *int, limit int, server_name, server_id string) string {
		var page int
		if offerset == nil {
			page = 0
		} else {
			page = *offerset
		}

		uri := fmt.Sprintf("https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/phones?offset=%d&limit=%d&server_name=%s&server_id=%s", page, limit, server_name, server_id)

		client := &http.Client{}
		req, _ := http.NewRequest("GET", uri, nil)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("X-Auth-Token", authtoken.Authtoken())
		resp, _ := client.Do(req)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("test ListCloudPhoneServers: ", string(body))
		return string(body)
	}

	if len(r.Form.Get("server_name")) > 0 {
		f(nil,0, r.Form.Get("server_name"), "")
	}
	if len(r.Form.Get("server_id")) > 0 {
		f(nil,0, "", r.Form.Get("server_id"))
	}

	page := 1
	limit := 100
	for i := 0; i < page; i++ {
		body := f(&page, limit, "", "")
		page++
		if len(body) == 0 {
			break
		}
	}

}
