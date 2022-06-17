package phoneinstance

import (
	"CPH-SDK/conf"
	"CPH-SDK/httphelper2"
	"CPH-SDK/response2"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type rcpBody struct {
	ImageId string `json:"image_id"`
	Phones  []struct {
		PhoneID  string `json:"phone_id"`
		Property string `json:"property"`
	} `json:"phones"`
}

// 重置云手机
func ResetCloudPhone(w http.ResponseWriter, r *http.Request) {
	resp := response2.NewResp()
	// var projectId string // 必填，项目ID
	// r.ParseForm()
	// if len(r.Form.Get("projectId")) > 0 {
	// 	projectId = r.Form.Get("projectId")
	// } else {
	// 	resp.BadReq(w)
	// 	return
	// }
	var rcp rcpBody
	err := json.NewDecoder(r.Body).Decode(&rcp)
	if err != nil {
		resp.BadReq(w)
		return
	}

	if len(rcp.Phones) == 0 {
		resp.BadReq(w)
		return
	}

	data, _ := json.Marshal(rcp)
	uri := fmt.Sprintf("https://%s/v1/%s/cloud-phone/phones/batch-reset", conf.Config.Huawei.Endpoint, conf.Config.Huawei.ProjectId)
	// uri := fmt.Sprintf("%s/%s/cloud-phone/phones/batch-reset", global.BaseUrl, projectId)
	body, err := httphelper2.HttpPost(uri, data)
	if err != nil {
		log.Println("ResetCloudPhone err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
