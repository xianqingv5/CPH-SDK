package keypair

import (
	"encoding/json"
	"fmt"
	"global"
	"httphelper"
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

type ukpBody struct {
	Servers []struct {
		KeypairName string `json:"keypair_name"`
		ServerID    string `json:"server_id"`
	} `json:"servers"`
}

func WriteTo(w http.ResponseWriter, data Res) {
	re, _ := json.Marshal(data)
	w.Write(re)
}

func UpdateKeypair(w http.ResponseWriter, r *http.Request) {
	var res Res
	var projectId string // 必填，项目ID
	var ukp ukpBody
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.status = requestErr
		WriteTo(w, res)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&ukp)
	if err != nil {
		res.status = requestErr
		WriteTo(w, res)
		return
	}

	if len(ukp.Servers) == 0 {
		res.status = requestErr
		WriteTo(w, res)
		return
	}
	mydata, _ := json.Marshal(ukp)

	uri := fmt.Sprintf("%s/%s/cloud-phone/servers/open-access", global.BaseUrl, projectId)
	body, err := httphelper.HttpPut(uri, mydata)
	res.status = OK
	if err != nil {
		res.status = requestErr
	} else {
		res.data = string(body)
	}
	WriteTo(w, res)
}
