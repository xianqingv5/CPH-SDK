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

type upnBody struct {
	PhoneName string `json:"phone_name"`
}

// 修改云手机名称
func UpdatePhoneName(w http.ResponseWriter, r *http.Request) {
	resp := response2.NewResp()

	// var projectId string // 必填，项目ID

	// r.ParseForm()
	// if len(r.Form.Get("projectId")) > 0 {
	// 	projectId = r.Form.Get("projectId")
	// } else {
	// 	resp.BadReq(w)
	// 	return
	// }
	var phoneId string // 必填，规格状态
	var upn upnBody
	if len(r.Form.Get("phone_id")) > 0 {
		phoneId = r.Form.Get("phone_id")
	} else {
		resp.BadReq(w)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&upn)
	if err != nil {
		resp.BadReq(w)
		return
	}

	if len(upn.PhoneName) == 0 {
		resp.BadReq(w)
		return
	}

	data, _ := json.Marshal(upn)
	uri := fmt.Sprintf("https://%s/v1/%s/cloud-phone/phone/phone_id=%s", conf.Config.Huawei.Endpoint, conf.Config.Huawei.ProjectId, phoneId)
	// uri := fmt.Sprintf("%s/%s/cloud-phone/phone/phone_id=%s", global.BaseUrl, projectId, phoneId)
	body, err := httphelper2.HttpPost(uri, data)
	if err != nil {
		log.Println("UpdatePhoneName err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
