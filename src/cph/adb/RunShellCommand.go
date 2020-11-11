package adb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"httphelper"
)

func RunShellCommand(w http.ResponseWriter, r *http.Request) {
	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/phones/commands"
	postbody := GetPostBody(r, "shell")
	if postbody == nil {
		return
	}

	d, _ := json.Marshal(postbody)

	body, err := httphelper.HttpPost(uri, d)
	if err != nil {
		return
	}
	fmt.Println("test RunShellCommand: ", string(body))

	WriteTo(w, body)
}