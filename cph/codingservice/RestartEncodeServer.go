package codingservice

import (
	"CPH-SDK/conf"
	"CPH-SDK/httphelper2"
	"CPH-SDK/response2"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type esIDs struct {
	EncodeServerIds []string `json:"encode_server_ids"`
}

func RestartEncodeServer(w http.ResponseWriter, r *http.Request) {
	resp := response2.NewResp()
	uri := fmt.Sprintf("https://%s/v1/%s/cloud-phone/encode-servers/batch-restart", conf.Config.Huawei.Endpoint, conf.Config.Huawei.ProjectId)
	// uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/encode-servers/batch-restart"

	var esids esIDs
	err := json.NewDecoder(r.Body).Decode(&esids)
	if err != nil {
		resp.BadReq(w)
		return
	}

	data, _ := json.Marshal(esids)
	body, err := httphelper2.HttpPost(uri, data)
	if err != nil {
		log.Println("RestartEncodeServer err: ", err)
		resp.IntervalServErr(w)
		return
	}
	fmt.Println("test RestartEncodeServer: ", string(body))

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
