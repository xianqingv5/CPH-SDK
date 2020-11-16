package task

import (
	"encoding/json"
	"fmt"
	"global"
	"httphelper"
	"net/http"
)

func QueryJobs(w http.ResponseWriter, r *http.Request) {
	var res Res
	var projectId string // 必填，项目ID
	var jobId string     // 必填，任务下发请求时响应的job_id
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	if len(r.Form.Get("jobId")) > 0 {
		jobId = r.Form.Get("jobId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}
	uri := fmt.Sprintf("%s/%s/cloud-phone/jobs/%s", global.BaseUrl, projectId, jobId)
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
