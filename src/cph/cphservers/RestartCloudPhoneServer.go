package cphservers

import (
	"encoding/json"
	"net/http"

	"httphelper"
)

type rcpsBody struct {
	ServerIDs []string `json:"server_ids"`
}

func RestartCloudPhoneServer(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST" {
		return
	}

	var rcps rcpsBody
	err := json.NewDecoder(r.Body).Decode(&rcps)
	if err != nil {
		return
	}

	if len(rcps.ServerIDs) == 0 {
		return
	}

	data, _ := json.Marshal(rcps)

	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/servers/batch-restart"
	body, err := httphelper.HttpPost(uri, data)
	if err != nil {
		return
	}

	WriteTo(w, body)
}