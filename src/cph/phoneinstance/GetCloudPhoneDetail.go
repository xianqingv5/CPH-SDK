package phoneinstance

import (
	"encoding/json"
	"fmt"
	"net/http"
	"httphelper"
	"global"
)

func GetCloudPhoneDetail (w http.ResponseWriter, r *http.Request) string {
	var res Res
	var projectId string // 必填，项目ID
	var phoneId string // 必填，云手机的唯一标识
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
	}

	if len(r.Form.Get("phone_id")) > 0 {
		phoneId = r.Form.Get("phone_id")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
	}
	uri := fmt.Sprintf("%s/%s/cloud-phone/phones/%s", global.BaseUrl, projectId, phoneId)
	body, _ := httphelper.HttpGet(uri)
	res.data = string(body)
	res.status = OK
	re, _ := json.Marshal(res)
	w.Write(re)
	return ""
}