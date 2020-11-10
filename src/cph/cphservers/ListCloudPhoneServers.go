package cphservers

import (
	"fmt"
	"httphelper"
	"net/http"
)

func ListCloudPhoneServers(w http.ResponseWriter, r *http.Request) {
	f := func(offset, limit, server_name, server_id string) ([]byte, error) {

		uri := fmt.Sprintf("https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/phones?offset=%s&limit=%s&server_name=%s&server_id=%s", offset, limit, server_name, server_id)

		body, err := httphelper.HttpGet(uri)
		fmt.Println("test ListCloudPhoneServers: ", string(body))
		return body, err
	}

	serverName := r.Form.Get("server_name")
	serverID := r.Form.Get("server_id")
	offset := r.Form.Get("offset")
	limit := r.Form.Get("limit")

	body, err := f(offset, limit, serverName, serverID)
	if err != nil {
		return
	}

	WriteTo(w, body)
}
