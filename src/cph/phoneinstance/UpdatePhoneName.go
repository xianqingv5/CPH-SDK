package phoneinstance

import (
	"encoding/json"
	"fmt"
	"net/http"
	"httphelper"
	"global"
)

func UpdatePhoneName (w http.ResponseWriter, r *http.Request) string {
	var res Res
	var projectId string // 必填，项目ID
	var phoneId string // 非必填，规格状态
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
	}

	if len(r.Form.Get("phone_id")) > 0 {
		phoneId = r.Form.Get("phone_id")
	}
	uri := fmt.Sprintf("%s/%s/cloud-phone/phone/phone_id=%s", global.BaseUrl, projectId, phoneId)
	body, _ := httphelper.HttpGet(uri)
	res.data = string(body)
	res.status = OK
	re, _ := json.Marshal(res)
	w.Write(re)
	return ""
}