package adb

import (
	"CPH-SDK/conf"
	"CPH-SDK/httphelper2"
	"CPH-SDK/response2"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// 执行异步adb shell命令
func RunShellCommand(w http.ResponseWriter, r *http.Request) {
	resp := response2.NewResp()
	uri := fmt.Sprintf("https://%s/v1/%s/cloud-phone/phones/commands", conf.Config.Huawei.Endpoint, conf.Config.Huawei.ProjectId)
	// uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/phones/commands"
	postbody := GetPostBody(r, "shell")
	if postbody == nil {
		resp.BadReq(w)
		return
	}

	d, _ := json.Marshal(postbody)

	body, err := httphelper2.HttpPost(uri, d)
	if err != nil {
		log.Println("RunShellCommand err: ", err)
		resp.IntervalServErr(w)
		return
	}
	fmt.Println("test RunShellCommand: ", string(body))

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
