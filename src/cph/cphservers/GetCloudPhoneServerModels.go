package cphservers

import (
	"net/http"
	"net/url"

	"httphelper"
	"util"
)

func GetCloudPhoneServerModels(w http.ResponseWriter, r *http.Request) {
	v := url.Values{}
	r.ParseForm()
	typeInfo := r.Form.Get("product_type")
	util.AddurlParam("product_type", typeInfo, &v)

	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/server-models"
	uri = uri + "?" + v.Encode()

	body, err := httphelper.HttpGet(uri)
	if err != nil {
		return
	}

	WriteTo(w, body)
}
