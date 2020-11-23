package cphservers

import (
	"encoding/json"
	"log"
	"net/http"

	"httphelper"
	"response"
)

type psfBody struct {
	BucketName string   `json:"bucket_name"`
	ObjectPath string   `json:"object_path"`
	ServerIDs  []string `json:"server_ids"`
}

// todo test
func PushShareFiles(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResp()

	if r.Method != "POST" {
		resp.BadReqMethod(w)
		return
	}

	var psf psfBody
	err := json.NewDecoder(r.Body).Decode(&psf)
	if err != nil {
		resp.BadReq(w)
		return
	}

	if len(psf.BucketName) == 0 || len(psf.ObjectPath) == 0 || len(psf.ServerIDs) == 0 {
		resp.BadReq(w)
		return
	}
	data, _ := json.Marshal(psf)

	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/phones/share-files"
	body, err := httphelper.HttpPost(uri, data)
	if err != nil {
		log.Println("PushShareFiles err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
