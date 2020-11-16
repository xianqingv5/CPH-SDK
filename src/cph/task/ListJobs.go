package task

import (
	// "authtoken"
	"encoding/json"
	"fmt"
	"global"
	"httphelper"
	// "io/ioutil"
	"net/http"
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

func ListJobs(w http.ResponseWriter, r *http.Request) {
	var res Res
	var projectId string // 必填，项目ID
	var requestId string // 必填，任务下发请求时响应的request_id
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	if len(r.Form.Get("requestId")) > 0 {
		requestId = r.Form.Get("requestId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}
	uri := fmt.Sprintf("%s/%s/cloud-phone/phones?request_id=%s", global.BaseUrl, projectId, requestId)
	//client := &http.Client{}
	//req, _ := http.NewRequest("GET", uri, nil)
	//req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("X-Auth-Token", authtoken.Authtoken())
	//resp, _ := client.Do(req)
	//body, _ := ioutil.ReadAll(resp.Body)
	body, err := httphelper.HttpGet(uri)
	res.status = OK
	if err != nil {
		res.status = requestErr
	} else {
		res.data = string(body)
	}
	re, _ := json.Marshal(res)

	w.Write(re)
}
