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

type pocpBody struct {
	PhoneIDs []string `json:"phone_ids"`
}

// 关闭云手机
func PowerOffCloudPhone(w http.ResponseWriter, r *http.Request) {
	resp := response2.NewResp()

	// var projectId string // 必填，项目ID

	// r.ParseForm()
	// if len(r.Form.Get("projectId")) > 0 {
	// 	projectId = r.Form.Get("projectId")
	// } else {
	// 	resp.BadReq(w)
	// 	return
	// }
	var pocp pocpBody
	err := json.NewDecoder(r.Body).Decode(&pocp)
	if err != nil {
		resp.BadReq(w)
		return
	}

	if len(pocp.PhoneIDs) == 0 {
		resp.BadReq(w)
		return
	}

	data, _ := json.Marshal(pocp)
	uri := fmt.Sprintf("https://%s/v1/%s/cloud-phone/batch-stop", conf.Config.Huawei.Endpoint, conf.Config.Huawei.ProjectId)
	// uri := fmt.Sprintf("%s/%s/cloud-phone/batch-stop", global.BaseUrl, projectId)
	body, err := httphelper2.HttpPost(uri, data)
	if err != nil {
		log.Println("PowerOffCloudPhone err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
