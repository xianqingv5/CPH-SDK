package phoneinstance

import (
	"encoding/json"
	"fmt"
	"global"
	"log"
	"net/http"

	"httphelper"
	"response"
)

type racpBody struct {
	Phones []struct {
		PhoneID  string `json:"phone_id"`
		Property string `json:"property"`
	} `json:"phones"`
	ImageID string `json:"image_id"`
}

func RestartCloudPhone(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResp()

	var projectId string // 必填，项目ID
	var racp racpBody
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		resp.BadReq(w)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&racp)
	if err != nil {
		resp.BadReq(w)
		return
	}

	if len(racp.Phones) == 0 {
		resp.BadReq(w)
		return
	}

	data, _ := json.Marshal(racp)
	uri := fmt.Sprintf("%s/%s/cloud-phone/phones/batch-restart", global.BaseUrl, projectId)
	body, err := httphelper.HttpPost(uri, data)
	if err != nil {
		log.Println("RestartCloudPhone err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
