package cphservers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"response"

	"httphelper"
	"util"
)

// 查询云手机服务器规格列表
func GetCloudPhoneServerModels(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResp()

	v := url.Values{}
	r.ParseForm()
	typeInfo := r.Form.Get("product_type")
	util.AddurlParam("product_type", typeInfo, &v)

	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/server-models"
	uri = uri + "?" + v.Encode()

	body, err := httphelper.HttpGet(uri)
	if err != nil {
		log.Println("GetCloudPhoneServerModels err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
