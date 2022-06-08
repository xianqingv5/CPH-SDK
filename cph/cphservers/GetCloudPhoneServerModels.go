package cphservers

import (
	"CPH-SDK/conf"
	"CPH-SDK/httphelper2"
	"CPH-SDK/response2"
	"CPH-SDK/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// 查询云手机服务器规格列表
func GetCloudPhoneServerModels(w http.ResponseWriter, r *http.Request) {
	resp := response2.NewResp()

	v := url.Values{}
	r.ParseForm()
	typeInfo := r.Form.Get("product_type")
	util.AddurlParam("product_type", typeInfo, &v)
	uri := fmt.Sprintf("https://%s/v1/%s/cloud-phone/server-models", conf.Config.Huawei.Endpoint, conf.Config.Huawei.ProjectId)
	// uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/server-models"
	uri = uri + "?" + v.Encode()

	body, err := httphelper2.HttpGet(uri)
	if err != nil {
		log.Println("GetCloudPhoneServerModels err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
