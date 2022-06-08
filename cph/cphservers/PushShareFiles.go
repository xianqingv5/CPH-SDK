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

type psfBody struct {
	BucketName string   `json:"bucket_name"`
	ObjectPath string   `json:"object_path"`
	ServerIDs  []string `json:"server_ids"`
}

// 推送共享存储文件
// todo test
func PushShareFiles(w http.ResponseWriter, r *http.Request) {
	resp := response2.NewResp()

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
	uri := fmt.Sprintf("https://%s/v1/%s/cloud-phone/phones/share-files", conf.Config.Huawei.Endpoint, conf.Config.Huawei.ProjectId)
	// uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/phones/share-files"
	body, err := httphelper2.HttpPost(uri, data)
	if err != nil {
		log.Println("PushShareFiles err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
