package phoneinstance

import (
	"encoding/json"
	"fmt"
	"global"
	"httphelper"
	"net/http"
)

func ListPhoneImages(w http.ResponseWriter, r *http.Request) {
	var res Res
	var projectId string // 必填，项目ID
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		res.status = requestErr
		re, _ := json.Marshal(res)
		w.Write(re)
		return
	}

	uri := fmt.Sprintf("%s/%s/cloud-phone/phone-images", global.BaseUrl, projectId)
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
