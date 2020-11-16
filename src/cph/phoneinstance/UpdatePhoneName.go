package phoneinstance

import (
	"encoding/json"
	"fmt"
	"global"
	"httphelper"
	"net/http"
)

type upnBody struct {
	PhoneName string `json:"phone_name"`
}

func UpdatePhoneName(w http.ResponseWriter, r *http.Request) {
	var res Res
	var projectId string // 必填，项目ID
	var phoneId string   // 必填，规格状态
	var upn upnBody
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

	err := json.NewDecoder(r.Body).Decode(&upn)
	if err != nil {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	if len(upn.PhoneName) == 0 {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	data, _ := json.Marshal(upn)
	uri := fmt.Sprintf("%s/%s/cloud-phone/phone/phone_id=%s", global.BaseUrl, projectId, phoneId)
	body, err := httphelper.HttpPost(uri, data)
	res.status = OK
	if err != nil {
		res.status = requestErr
	} else {
		res.data = string(body)
	}
	re, _ := json.Marshal(res)
	w.Write(re)
}
