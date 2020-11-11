package cphservers

import (
	"encoding/json"
	"net/http"

	"httphelper"
)

type postBody struct {
	FilePaths string   `json:"file_paths"`
	ServerIds []string `json:"server_ids"`
}

// todo test
func DeleteShareFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	var pb postBody
	if err := json.NewDecoder(r.Body).Decode(&pb); err != nil {
		return
	}

	if len(pb.FilePaths) == 0 || len(pb.ServerIds) == 0 {
		return
	}

	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/phones/share-files"
	data, _ := json.Marshal(pb)

	body, err := httphelper.HttpPost(uri, data)
	if err != nil {
		return
	}

	WriteTo(w, body)
}
