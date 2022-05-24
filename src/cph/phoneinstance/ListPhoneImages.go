package phoneinstance

import (
	"encoding/json"
	"fmt"
	"global"
	"log"
	"net/http"

	"httphelper"
	"response"
)

// 查询手机镜像
func ListPhoneImages(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResp()

	var projectId string // 必填，项目ID
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		resp.BadReq(w)
		return
	}

	uri := fmt.Sprintf("%s/%s/cloud-phone/phone-images", global.BaseUrl, projectId)
	body, err := httphelper.HttpGet(uri)
	if err != nil {
		log.Println("ListPhoneImages err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
