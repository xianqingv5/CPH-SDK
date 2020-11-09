package task

import (
	"authtoken"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func QueryJobs (w http.ResponseWriter, r *http.Request) string {
	var res Res
	var projectId string // 必填，项目ID
	var jobId string // 必填，任务下发请求时响应的job_id
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
	}

	if len(r.Form.Get("jobId")) > 0 {
		jobId = r.Form.Get("jobId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
	}
	uri := fmt.Sprintf("https://cph.cn-east-3.myhuaweicloud.com/v1/%s/cloud-phone/jobs/%s", projectId, jobId)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", authtoken.Authtoken())
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	res.data = string(body)
	res.status = OK
	re, _ := json.Marshal(res)
	w.Write(re)
}