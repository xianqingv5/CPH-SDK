package phoneinstance

import (
	"CPH-SDK/conf"
	"CPH-SDK/httphelper2"
	"CPH-SDK/response2"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"CPH-SDK/util"
)

// 查询云手机列表
func ListCloudPhones(w http.ResponseWriter, r *http.Request) {
	f := func(offset, limit, phone_name, server_id, status, typeInfo string) ([]byte, error) {
		v := url.Values{}
		util.AddurlParam("offset", offset, &v)
		util.AddurlParam("limit", limit, &v)
		util.AddurlParam("server_name", phone_name, &v)
		util.AddurlParam("server_id", server_id, &v)
		util.AddurlParam("status", status, &v)
		util.AddurlParam("type", typeInfo, &v)

		// uri := fmt.Sprintf("%s/%s/cloud-phone/phones", global.BaseUrl, projectId)
		uri := fmt.Sprintf("https://%s/v1/%s/cloud-phone/phones", conf.Config.Huawei.Endpoint, conf.Config.Huawei.ProjectId)
		uri = uri + "?" + v.Encode()
		body, err := httphelper2.HttpGet(uri)
		return body, err
	}

	resp := response2.NewResp()

	// var projectId string
	// r.ParseForm()
	// if len(r.Form.Get("projectId")) > 0 {
	// 	projectId = r.Form.Get("projectId")
	// } else {
	// 	resp.BadReq(w)
	// 	return
	// }
	fmt.Println("[test3]")
	phoneName := r.Form.Get("phone_name")
	serverID := r.Form.Get("server_id")
	offset := r.Form.Get("offset")
	limit := r.Form.Get("limit")
	status := r.Form.Get("status")
	typeInfo := r.Form.Get("type")
	fmt.Println("[test]")
	// body, err := f(offset, limit, phoneName, serverID, status, typeInfo, projectId)
	body, err := f(offset, limit, phoneName, serverID, status, typeInfo)
	if err != nil {
		log.Println("ListCloudPhones err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
