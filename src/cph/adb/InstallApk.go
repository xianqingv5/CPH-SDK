package adb

import (
	"encoding/json"
	"fmt"
	"httphelper"
	"log"
	"net/http"
	"response"
)

type AdbPostBody struct {
	Command   string   `json:"command"`
	Content   string   `json:"content"`
	ServerIds []string `json:"server_ids,omitempty"`
	PhoneIds  []string `json:"phone_ids,omitempty"`
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
	resp := response.NewResp()

	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/phones/commands"
	postbody := GetPostBody(r, "install")
	if postbody == nil {
		resp.BadReq(w)
		return
	}

	d, _ := json.Marshal(postbody)

	body, err := httphelper.HttpPost(uri, d)
	if err != nil {
		log.Println("InstallApk err: ", err)
		resp.IntervalServErr(w)
		return
	}
	fmt.Println("test InstallApk: ", string(body))

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
