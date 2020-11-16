package phoneinstance

import (
	"encoding/json"
	"fmt"
	"global"
	"httphelper"
	"net/http"
)

type rcpBody struct {
	Phones []struct {
		PhoneID  string `json:"phone_id"`
		Property string `json:"property"`
	} `json:"phones"`
}

func ResetCloudPhone(w http.ResponseWriter, r *http.Request) {
	var res Res
	var projectId string // 必填，项目ID
	var rcp rcpBody
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&rcp)
	if err != nil {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	if len(rcp.Phones) == 0 {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	data, _ := json.Marshal(rcp)
	uri := fmt.Sprintf("%s/%s/cloud-phone/phones/batch-reset", global.BaseUrl, projectId)
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
