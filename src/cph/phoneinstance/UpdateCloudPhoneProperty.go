package phoneinstance

import (
	"encoding/json"
	"fmt"
	"global"
	"httphelper"
	"log"
	"net/http"
	"response"
)

type ucpBody struct {
	Phones []struct {
		PhoneID  string `json:"phone_id"`
		Property string `json:"property"`
	} `json:"phones"`
}

// 更新云手机属性
func UpdateCloudPhoneProperty(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResp()

	var projectId string // 必填，项目ID
	var ucp ucpBody
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		resp.BadReq(w)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&ucp)
	if err != nil {
		resp.BadReq(w)
		return
	}

	if len(ucp.Phones) == 0 {
		resp.BadReq(w)
		return
	}

	data, _ := json.Marshal(ucp)
	uri := fmt.Sprintf("%s/%s/cloud-phone/phones/batch-update-property", global.BaseUrl, projectId)
	body, err := httphelper.HttpPost(uri, data)
	if err != nil {
		log.Println("UpdateCloudPhoneProperty err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
