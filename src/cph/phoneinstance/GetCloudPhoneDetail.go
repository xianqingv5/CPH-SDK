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

// 查询云手机详情
func GetCloudPhoneDetail(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResp()

	var projectId string // 必填，项目ID
	var phoneId string   // 必填，云手机的唯一标识
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		resp.BadReq(w)
		return
	}

	if len(r.Form.Get("phone_id")) > 0 {
		phoneId = r.Form.Get("phone_id")
	} else {
		resp.BadReq(w)
		return
	}
	uri := fmt.Sprintf("%s/%s/cloud-phone/phones/%s", global.BaseUrl, projectId, phoneId)
	body, err := httphelper.HttpGet(uri)
	if err != nil {
		log.Println("GetCloudPhoneDetail err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
