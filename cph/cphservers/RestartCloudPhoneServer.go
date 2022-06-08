package cphservers

import (
	"CPH-SDK/conf"
	"CPH-SDK/httphelper2"
	"CPH-SDK/response2"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type rcpsBody struct {
	ServerIDs []string `json:"server_ids"`
}

// 重启云手机
func RestartCloudPhoneServer(w http.ResponseWriter, r *http.Request) {
	resp := response2.NewResp()

	if r.Method != "POST" {
		resp.BadReqMethod(w)
		return
	}

	var rcps rcpsBody
	err := json.NewDecoder(r.Body).Decode(&rcps)
	if err != nil {
		resp.BadReq(w)
		return
	}

	if len(rcps.ServerIDs) == 0 {
		resp.BadReq(w)
		return
	}

	data, _ := json.Marshal(rcps)
	uri := fmt.Sprintf("https://%s/v1/%s/cloud-phone/servers/batch-restart", conf.Config.Huawei.Endpoint, conf.Config.Huawei.ProjectId)
	// uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/servers/batch-restart"
	body, err := httphelper2.HttpPost(uri, data)
	if err != nil {
		log.Println("RestartCloudPhoneServer err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
