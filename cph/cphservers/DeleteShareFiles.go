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

type postBody struct {
	FilePaths string   `json:"file_paths"`
	ServerIds []string `json:"server_ids"`
}

// 删除共享存储文件
// todo test
func DeleteShareFiles(w http.ResponseWriter, r *http.Request) {
	resp := response2.NewResp()

	if r.Method != "POST" {
		resp.BadReqMethod(w)
		return
	}

	var pb postBody
	if err := json.NewDecoder(r.Body).Decode(&pb); err != nil {
		resp.BadReq(w)
		return
	}

	if len(pb.FilePaths) == 0 || len(pb.ServerIds) == 0 {
		resp.BadReq(w)
		return
	}
	uri := fmt.Sprintf("https://%s/v1/%s/cloud-phone/phones/share-files", conf.Config.Huawei.Endpoint, conf.Config.Huawei.ProjectId)
	// uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/phones/share-files"
	data, _ := json.Marshal(pb)

	body, err := httphelper2.HttpPost(uri, data)
	if err != nil {
		log.Println("DeleteShareFiles err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
