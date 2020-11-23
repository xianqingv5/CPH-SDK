package task

import (
	// "authtoken"
	"encoding/json"
	"fmt"
	"global"
	"log"
	"response"

	"httphelper"
	"net/http"
)

func ListJobs(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResp()

	var projectId string // 必填，项目ID
	var requestId string // 必填，任务下发请求时响应的request_id
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		resp.BadReq(w)
		return
	}

	if len(r.Form.Get("requestId")) > 0 {
		requestId = r.Form.Get("requestId")
	} else {
		resp.BadReq(w)
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
	if err != nil {
		log.Println("ListJobs err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
