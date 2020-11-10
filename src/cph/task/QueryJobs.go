package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"httphelper"
	"global"
)

func QueryJobs (w http.ResponseWriter, r *http.Request) string {
	var res global.Res
	var projectId string // 必填，项目ID
	var jobId string // 必填，任务下发请求时响应的job_id
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.Status = global.StatusBadRequest
		re, _ := json.Marshal(res)
		w.Write(re)
	}

	if len(r.Form.Get("jobId")) > 0 {
		jobId = r.Form.Get("jobId")
	} else {
		res.Status = global.StatusBadRequest
		re, _ := json.Marshal(res)
		w.Write(re)
	}
	uri := fmt.Sprintf("%s/%s/cloud-phone/jobs/%s", global.BaseUrl, projectId, jobId)
	body, _ := httphelper.HttpGet(uri)
	res.Data = string(body)
	res.Status = global.StatusOK
	res.Info = global.StatusText(global.StatusOK)
	re, _ := json.Marshal(res)
	w.Write(re)
	return ""
}