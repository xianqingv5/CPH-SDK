package task

import (
	"CPH-SDK/conf"
	"CPH-SDK/httphelper2"
	"CPH-SDK/response2"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// 查询任务执行状态
func QueryJobs(w http.ResponseWriter, r *http.Request) {
	resp := response2.NewResp()
	// var projectId string // 必填，项目ID

	// if len(r.Form.Get("projectId")) > 0 {
	// 	projectId = r.Form.Get("projectId")
	// } else {
	// 	resp.BadReq(w)
	// 	return
	// }
	var jobId string // 必填，任务下发请求时响应的job_id
	if len(r.URL.Query().Get("jobId")) > 0 {
		jobId = r.URL.Query().Get("jobId")
	} else {
		resp.BadReq(w)
		return
	}

	uri := fmt.Sprintf("https://%s/v1/%s/cloud-phone/jobs/%s", conf.Config.Huawei.Endpoint, conf.Config.Huawei.ProjectId, jobId)
	// uri := "https://cph.ru-northwest-2.myhuaweicloud.com/v1/0fa76c3cc35341559eea89ca551cae88/cloud-phone/jobs/c0a518442a0141d18164e463aaba41d3"
	body, err := httphelper2.HttpGet(uri)
	if err != nil {
		log.Println("QueryJobs err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
