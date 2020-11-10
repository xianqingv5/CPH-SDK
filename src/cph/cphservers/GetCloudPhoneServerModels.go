package cphservers

import (
	"fmt"
	"httphelper"
	"net/http"
)

func GetCloudPhoneServerModels(w http.ResponseWriter, r *http.Request)  {
	typeInfo := r.Form.Get("type")

	uri := fmt.Sprintf("https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/server-models?product_type=%s", typeInfo)

	body, err := httphelper.HttpGet(uri)
	if err != nil {
		return
	}

	WriteTo(w, body)
}