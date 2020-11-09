package phoneinstance

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
	"authtoken"
)

// 数据返回格式
type Res struct {
	status int
	data string
}

// 状态码
const (
	OK 				= 200 // 成功
	requestErr 		= 400 // 客户端错误
)

func GetCloudPhoneDetail(w http.ResponseWriter, r *http.Request) {
	f := func(offerset *int, limit int, server_name, server_id string, projectId string) string {
		var page int
		if offerset == nil {
			page = 0
		} else {
			page = *offerset
		}

		uri := fmt.Sprintf("https://cph.cn-east-3.myhuaweicloud.com/v1/%s/cloud-phone/phones?offset=%d&limit=%d&server_name=%s&server_id=%s",projectId, page, limit, server_name, server_id)

		client := &http.Client{}
		req, _ := http.NewRequest("GET", uri, nil)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("X-Auth-Token", authtoken.Authtoken())
		resp, _ := client.Do(req)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("test ListCloudPhoneServers: ", string(body))
		return string(body)
	}

	var projectId string
	var res Res
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
	}

	if len(r.Form.Get("server_name")) > 0 {
		f(nil,0, r.Form.Get("server_name"), "", projectId)
	}
	if len(r.Form.Get("server_id")) > 0 {
		f(nil,0, "", r.Form.Get("server_id"), projectId)
	}

	page := 1
	limit := 100
	for i := 0; i < page; i++ {
		body := f(&page, limit, "", "", projectId)
		page++
		if len(body) == 0 {
			break
		}
	}

}
