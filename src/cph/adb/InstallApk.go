package adb

import (
	"encoding/json"
	"fmt"
	"httphelper"
	"net/http"
)

type AdbPostBody struct {
	Command   string   `json:"command"`
	Content   string   `json:"content"`
	ServerIds []string `json:"server_ids,omitempty"`
	PhoneIds  []string `json:"phone_ids,omitempty"`
}

func WriteTo(w http.ResponseWriter, data []byte) {
	w.Write(data)
}

func GetPostBody(r *http.Request, format string) *AdbPostBody {
	var adb AdbPostBody

	err := json.NewDecoder(r.Body).Decode(&adb)
	if err != nil {
		return nil
	}

	if (len(adb.ServerIds) == 0 && len(adb.PhoneIds) == 0) || len(adb.Content) == 0 {
		return nil
	}

	adb.Command = format

	return &adb
}

func InstallApk(w http.ResponseWriter, r *http.Request) {
	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/phones/commands"
	postbody := GetPostBody(r, "install")
	if postbody == nil {
		return
	}

	d, _ := json.Marshal(postbody)

	body, err := httphelper.HttpPost(uri, d)
	if err != nil {
		return
	}
	fmt.Println("test InstallApk: ", string(body))

	WriteTo(w, body)
}
