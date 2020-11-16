package phoneinstance

import (
	"encoding/json"
	"fmt"
	"global"
	"httphelper"
	"net/http"
)

type racpBody struct {
	Phones []struct {
		PhoneID  string `json:"phone_id"`
		Property string `json:"property"`
	} `json:"phones"`
	ImageID string `json:"image_id"`
}

func RestartCloudPhone(w http.ResponseWriter, r *http.Request) {
	var res Res
	var projectId string // 必填，项目ID
	var racp racpBody
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&racp)
	if err != nil {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	if len(racp.Phones) == 0 {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	data, _ := json.Marshal(racp)
	uri := fmt.Sprintf("%s/%s/cloud-phone/phones/batch-restart", global.BaseUrl, projectId)
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
