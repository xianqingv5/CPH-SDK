package task

import (
	"encoding/json"
	"fmt"
	"global"
	"log"
	"net/http"

	"httphelper"
	"response"
)

// 查询任务执行状态
func QueryJobs(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResp()

	var projectId string // 必填，项目ID
	var jobId string     // 必填，任务下发请求时响应的job_id
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		resp.BadReq(w)
		return
	}

	if len(r.Form.Get("jobId")) > 0 {
		jobId = r.Form.Get("jobId")
	} else {
		resp.BadReq(w)
		return
	}
	uri := fmt.Sprintf("%s/%s/cloud-phone/jobs/%s", global.BaseUrl, projectId, jobId)
	body, err := httphelper.HttpGet(uri)
	if err != nil {
		log.Println("QueryJobs err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
