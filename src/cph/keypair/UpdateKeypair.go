package keypair

import (
	"encoding/json"
	"fmt"
	"net/http"
	"httphelper"
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

func UpdateKeypair (w http.ResponseWriter, r *http.Request) string {
	var res Res
	var projectId string // 必填，项目ID
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
	}

	uri := fmt.Sprintf("%s/%s/cloud-phone/servers/open-access", global.BaseUrl, projectId)
	body, _ := httphelper.HttpGet(uri)
	res.data = string(body)
	res.status = OK
	re, _ := json.Marshal(res)
	w.Write(re)
	return ""
}