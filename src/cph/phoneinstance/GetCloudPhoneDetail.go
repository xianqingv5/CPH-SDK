package phoneinstance

import (
	"encoding/json"
	"fmt"
	"global"
	"httphelper"
	"net/http"
)

func GetCloudPhoneDetail(w http.ResponseWriter, r *http.Request) {
	var res Res
	var projectId string // 必填，项目ID
	var phoneId string   // 必填，云手机的唯一标识
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	if len(r.Form.Get("phone_id")) > 0 {
		phoneId = r.Form.Get("phone_id")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}
	uri := fmt.Sprintf("%s/%s/cloud-phone/phones/%s", global.BaseUrl, projectId, phoneId)
	body, err := httphelper.HttpGet(uri)
	res.status = OK
	if err != nil {
		res.status = requestErr
	} else {
		res.data = string(body)
	}
	re, _ := json.Marshal(res)
	w.Write(re)

}
