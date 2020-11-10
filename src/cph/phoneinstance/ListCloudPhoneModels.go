package phoneinstance

import (
	"encoding/json"
	"fmt"
	"net/http"
	"httphelper"
	"global"
)

func ListCloudPhoneModels (w http.ResponseWriter, r *http.Request) string {
	var res Res
	var projectId string // 必填，项目ID
	var status string // 非必填，规格状态
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
	}

	if len(r.Form.Get("status")) > 0 {
		status = r.Form.Get("status")
	}
	uri := fmt.Sprintf("%s/%s/cloud-phone/phone-models/status=%s", global.BaseUrl, projectId, status)
	body, _ := httphelper.HttpGet(uri)
	res.data = string(body)
	res.status = OK
	re, _ := json.Marshal(res)
	w.Write(re)
	return ""
}