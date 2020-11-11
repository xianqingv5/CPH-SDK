package cphservers

import (
	"encoding/json"
	"httphelper"
	"net/http"
)

type psfBody struct {
	BucketName string   `json:"bucket_name"`
	ObjectPath string   `json:"object_path"`
	ServerIDs  []string `json:"server_ids"`
}

// todo test
func PushShareFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	var psf psfBody
	err := json.NewDecoder(r.Body).Decode(&psf)
	if err != nil {
		return
	}

	if len(psf.BucketName) == 0 || len(psf.ObjectPath) == 0 || len(psf.ServerIDs) == 0 {
		return
	}
	data, _ := json.Marshal(psf)

	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/phones/share-files"
	body, err := httphelper.HttpPost(uri, data)
	if err != nil {
		return
	}

	WriteTo(w, body)
}
