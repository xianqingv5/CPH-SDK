package adb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"authtoken"
)

type PostBody struct {
	Command   string   `json:"command"`
	Content   string   `json:"content"`
	ServerIds []string `json:"server_ids,omitempty"`
	PhoneIds  []string `json:"phone_ids,omitempty"`
}

func WriteTo(w http.ResponseWriter, data []byte) {
	w.Write(data)
}

func GetPostBody(r *http.Request, format string) *PostBody {
	sids := r.Form.Get("server_ids")
	s := strings.Split(sids, ",")
	pids := r.Form.Get("phone_ids")
	p := strings.Split(pids, ",")
	content := r.Form.Get("content")

	if (len(sids) == 0 && len(pids) == 0) || len(content) == 0 {
		return nil
	}

	return &PostBody{
		Command:   format,
		Content:   content,
		ServerIds: s,
		PhoneIds:  p,
	}
}

func InstallApk(w http.ResponseWriter, r *http.Request) {
	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/phones/commands"
	postbody := GetPostBody(r, "install")
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
	fmt.Println("test InstallApk: ", string(body))

	WriteTo(w, body)
}
