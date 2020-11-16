package phoneinstance

import (
	"encoding/json"
	"fmt"
	"global"
	"httphelper"
	"net/http"
	"net/url"
	"util"
)

func ListCloudPhoneModels(w http.ResponseWriter, r *http.Request) {
	var res Res
	var projectId string // 必填，项目ID
	var status string    // 非必填，规格状态
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	v := url.Values{}
	status = r.Form.Get("status")
	util.AddurlParam("status", status, &v)

	uri := fmt.Sprintf("%s/%s/cloud-phone/phone-models", global.BaseUrl, projectId)
	uri = uri + "?" + v.Encode()

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
