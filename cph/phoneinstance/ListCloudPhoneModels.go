package phoneinstance

import (
	"CPH-SDK/conf"
	"CPH-SDK/httphelper2"
	"CPH-SDK/response2"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"CPH-SDK/util"
)

// 查询云手机规格列表
func ListCloudPhoneModels(w http.ResponseWriter, r *http.Request) {
	resp := response2.NewResp()

	// var projectId string // 必填，项目ID

	// r.ParseForm()
	// if len(r.Form.Get("projectId")) > 0 {
	// 	projectId = r.Form.Get("projectId")
	// } else {
	// 	resp.BadReq(w)
	// 	return
	// }
	var status string // 非必填，规格状态
	v := url.Values{}
	status = r.Form.Get("status")
	util.AddurlParam("status", status, &v)
	uri := fmt.Sprintf("https://%s/v1/%s/cloud-phone/phone-models", conf.Config.Huawei.Endpoint, conf.Config.Huawei.ProjectId)
	// uri := fmt.Sprintf("%s/%s/cloud-phone/phone-models", global.BaseUrl, projectId)
	uri = uri + "?" + v.Encode()

	body, err := httphelper2.HttpGet(uri)
	if err != nil {
		log.Println("ListCloudPhoneModels err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
