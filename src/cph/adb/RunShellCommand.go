package adb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"httphelper"
	"response"
)

// 执行异步adb shell命令
func RunShellCommand(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResp()

	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/phones/commands"
	postbody := GetPostBody(r, "shell")
	if postbody == nil {
		resp.BadReq(w)
		return
	}

	d, _ := json.Marshal(postbody)

	body, err := httphelper.HttpPost(uri, d)
	if err != nil {
		log.Println("RunShellCommand err: ", err)
		resp.IntervalServErr(w)
		return
	}
	fmt.Println("test RunShellCommand: ", string(body))

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
