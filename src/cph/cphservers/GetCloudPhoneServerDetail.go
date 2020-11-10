package cphservers

import (
	"httphelper"
	"net/http"
)

func GetCloudPhoneServerDetail(w http.ResponseWriter, r *http.Request)  {
	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/servers/88b8449b896f3a4f0ad57222dd91909"

	body, err := httphelper.HttpGet(uri)
	if err != nil {
		return
	}

	WriteTo(w, body)
}