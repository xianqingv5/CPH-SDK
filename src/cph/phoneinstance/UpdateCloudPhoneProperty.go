package phoneinstance

import (
	"encoding/json"
	"fmt"
	"global"
	"httphelper"
	"net/http"
)

type ucpBody struct {
	Phones []struct {
		PhoneID  string `json:"phone_id"`
		Property string `json:"property"`
	} `json:"phones"`
}

func UpdateCloudPhoneProperty(w http.ResponseWriter, r *http.Request) {
	var res Res
	var projectId string // 必填，项目ID
	var ucp ucpBody
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&ucp)
	if err != nil {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	if len(ucp.Phones) == 0 {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	data, _ := json.Marshal(ucp)
	uri := fmt.Sprintf("%s/%s/cloud-phone/phones/batch-update-property", global.BaseUrl, projectId)
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
