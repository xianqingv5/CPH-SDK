package cphservers

import (
	"net/http"

	"httphelper"
)

// todo test
func GetCloudPhoneServerDetail(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	serverID := r.Form.Get("server_id")
	if len(serverID) == 0 {
		return
	}

	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/servers/" + serverID

	body, err := httphelper.HttpGet(uri)
	if err != nil {
		return
	}

	WriteTo(w, body)
}