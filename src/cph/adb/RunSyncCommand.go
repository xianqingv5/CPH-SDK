package adb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"authtoken"
)

func RunSyncCommand(w http.ResponseWriter, r *http.Request) {
	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/phones/sync-commands"
	postbody := GetPostBody(r, "shell")
	if postbody == nil {
		return
	}

	d, _ := json.Marshal(postbody)
	data := bytes.NewReader(d)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", uri, data)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", authtoken.Authtoken())
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("test RunSyncCommand: ", string(body))

	WriteTo(w, body)
}