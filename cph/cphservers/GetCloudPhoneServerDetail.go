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

// 查询云手机服务器详情
// todo test
func GetCloudPhoneServerDetail(w http.ResponseWriter, r *http.Request) {
	resp := response2.NewResp()

	r.ParseForm()
	serverID := r.Form.Get("server_id")
	if len(serverID) == 0 {
		resp.BadReq(w)
		return
	}
	uri := fmt.Sprintf("https://%s/v1/%s/cloud-phone/servers/%s", conf.Config.Huawei.Endpoint, conf.Config.Huawei.ProjectId, serverID)
	// uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/servers/" + serverID

	body, err := httphelper2.HttpGet(uri)
	if err != nil {
		log.Println("GetCloudPhoneServerDetail err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
