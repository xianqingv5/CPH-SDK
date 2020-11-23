package cphservers

import (
	"encoding/json"
	"log"
	"net/http"

	"httphelper"
	"response"
)

// todo test
func GetCloudPhoneServerDetail(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResp()

	r.ParseForm()
	serverID := r.Form.Get("server_id")
	if len(serverID) == 0 {
		resp.BadReq(w)
		return
	}

	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/servers/" + serverID

	body, err := httphelper.HttpGet(uri)
	if err != nil {
		log.Println("GetCloudPhoneServerDetail err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
