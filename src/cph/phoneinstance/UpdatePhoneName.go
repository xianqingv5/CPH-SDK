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

type upnBody struct {
	PhoneName string `json:"phone_name"`
}

func UpdatePhoneName(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResp()

	var projectId string // 必填，项目ID
	var phoneId string   // 必填，规格状态
	var upn upnBody
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		resp.BadReq(w)
		return
	}

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
	uri := fmt.Sprintf("%s/%s/cloud-phone/phone/phone_id=%s", global.BaseUrl, projectId, phoneId)
	body, err := httphelper.HttpPost(uri, data)
	if err != nil {
		log.Println("UpdatePhoneName err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
