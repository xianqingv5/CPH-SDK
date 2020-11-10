package phoneinstance

import (
	"encoding/json"
	"fmt"
	"httphelper"
	"net/http"
	"global"
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

func ListCloudPhones(w http.ResponseWriter, r *http.Request) {
	f := func(offerset *int, limit int, phone_name, server_id string, projectId string) string {
		var page int
		if offerset == nil {
			page = 0
		} else {
			page = *offerset
		}

		uri := fmt.Sprintf("%s/%s/cloud-phone/phones?offset=%d&limit=%d&server_name=%s&server_id=%s",projectId, page, limit, phone_name, server_id)
		body, _ := httphelper.HttpGet(uri)
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

	if len(r.Form.Get("phone_name")) > 0 {
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