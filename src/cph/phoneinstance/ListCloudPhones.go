package phoneinstance

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"response"

	"global"
	"httphelper"
	"util"
)

// 查询云手机列表
func ListCloudPhones(w http.ResponseWriter, r *http.Request) {
	f := func(offset, limit, phone_name, server_id, status, typeInfo string, projectId string) ([]byte, error) {
		v := url.Values{}
		util.AddurlParam("offset", offset, &v)
		util.AddurlParam("limit", limit, &v)
		util.AddurlParam("server_name", phone_name, &v)
		util.AddurlParam("server_id", server_id, &v)
		util.AddurlParam("status", status, &v)
		util.AddurlParam("type", typeInfo, &v)

		uri := fmt.Sprintf("%s/%s/cloud-phone/phones", global.BaseUrl, projectId)
		uri = uri + "?" + v.Encode()

		body, err := httphelper.HttpGet(uri)
		return body, err
	}

	resp := response.NewResp()

	var projectId string
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		resp.BadReq(w)
		return
	}

	phoneName := r.Form.Get("phone_name")
	serverID := r.Form.Get("server_id")
	offset := r.Form.Get("offset")
	limit := r.Form.Get("limit")
	status := r.Form.Get("status")
	typeInfo := r.Form.Get("type")

	body, err := f(offset, limit, phoneName, serverID, status, typeInfo, projectId)
	if err != nil {
		log.Println("ListCloudPhones err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
