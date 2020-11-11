package codingservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"httphelper"
)

type esIDs struct {
	EncodeServerIds []string `json:"encode_server_ids"`
}

func RestartEncodeServer(w http.ResponseWriter, r *http.Request) {
	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/encode-servers/batch-restart"

	var esids esIDs
	err := json.NewDecoder(r.Body).Decode(&esids)
	if err != nil {
		return
	}

	data, _ := json.Marshal(esids)
	body, err := httphelper.HttpPost(uri, data)
	if err != nil {
		return
	}
	fmt.Println("test RestartEncodeServer: ", string(body))

	WriteTo(w, body)
}
