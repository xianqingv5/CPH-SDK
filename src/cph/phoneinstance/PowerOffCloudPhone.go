package phoneinstance

import (
	"encoding/json"
	"fmt"
	"global"
	"httphelper"
	"net/http"
)

type pocpBody struct {
	PhoneIDs []string `json:"phone_ids"`
}

func PowerOffCloudPhone(w http.ResponseWriter, r *http.Request) {
	var res Res
	var projectId string // 必填，项目ID
	var pocp pocpBody
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&pocp)
	if err != nil {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	if len(pocp.PhoneIDs) == 0 {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	data, _ := json.Marshal(pocp)
	uri := fmt.Sprintf("%s/%s/cloud-phone/batch-stop", global.BaseUrl, projectId)
	body, err := httphelper.HttpPost(uri, data)
	res.status = OK
	if err != nil {
		res.status = requestErr
	} else {
		res.data = string(body)
	}
	re, _ := json.Marshal(res)
	w.Write(re)

	return
}
