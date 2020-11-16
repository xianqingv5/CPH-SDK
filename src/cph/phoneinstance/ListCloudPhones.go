package phoneinstance

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"global"
	"httphelper"
	"util"
)

// 数据返回格式
type Res struct {
	status int
	data   string
}

// 状态码
const (
	OK         = 200 // 成功
	requestErr = 400 // 客户端错误
)

func ListCloudPhones(w http.ResponseWriter, r *http.Request) {
	f := func(offset, limit, phone_name, server_id, status, typeInfo string, projectId string) ([]byte, error) {
		v := url.Values{}
		util.AddurlParam("offset", offset, &v)
		util.AddurlParam("limit", limit, &v)
		util.AddurlParam("server_name", phone_name, &v)
		util.AddurlParam("server_id", server_id, &v)
		util.AddurlParam("status", status, &v)
		util.AddurlParam("type", typeInfo, &v)

		uri := fmt.Sprintf("%s/%s/cloud-phone/phones", global.BaseUrl, projectId)
		uri = uri + "?" + v.Encode()

		body, err := httphelper.HttpGet(uri)
		return body, err
	}

	var projectId string
	var res Res
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	phoneName := r.Form.Get("phone_name")
	serverID := r.Form.Get("server_id")
	offset := r.Form.Get("offset")
	limit := r.Form.Get("limit")
	status := r.Form.Get("status")
	typeInfo := r.Form.Get("type")

	data, err := f(offset, limit, phoneName, serverID, status, typeInfo, projectId)
	res.status = OK
	if err != nil {
		res.status = requestErr
	} else {
		res.data = string(data)
	}

	myData, _ := json.Marshal(res)

	w.Write(myData)
}
