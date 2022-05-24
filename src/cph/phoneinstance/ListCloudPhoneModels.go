package phoneinstance

import (
	"encoding/json"
	"fmt"
	"global"
	"log"
	"net/http"
	"net/url"

	"httphelper"
	"response"
	"util"
)

// 查询云手机规格列表
func ListCloudPhoneModels(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResp()

	var projectId string // 必填，项目ID
	var status string    // 非必填，规格状态
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		resp.BadReq(w)
		return
	}

	v := url.Values{}
	status = r.Form.Get("status")
	util.AddurlParam("status", status, &v)

	uri := fmt.Sprintf("%s/%s/cloud-phone/phone-models", global.BaseUrl, projectId)
	uri = uri + "?" + v.Encode()

	body, err := httphelper.HttpGet(uri)
	if err != nil {
		log.Println("ListCloudPhoneModels err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
