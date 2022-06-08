package codingservice

import (
	"CPH-SDK/conf"
	"CPH-SDK/httphelper2"
	"CPH-SDK/response2"
	"CPH-SDK/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// 查询编码服务
func ListEncodeService(w http.ResponseWriter, r *http.Request) {
	f := func(offset, limit, types, status, server_id string) ([]byte, error) {
		v := url.Values{}
		util.AddurlParam("offset", offset, &v)
		util.AddurlParam("limit", limit, &v)
		util.AddurlParam("type", types, &v)
		util.AddurlParam("status", status, &v)
		util.AddurlParam("server_id", server_id, &v)
		uri := fmt.Sprintf("https://%s/v1/%s/cloud-phone/encode-servers", conf.Config.Huawei.Endpoint, conf.Config.Huawei.ProjectId)
		// uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/encode-servers"
		uri = uri + "?" + v.Encode()

		body, err := httphelper2.HttpGet(uri)
		if err != nil {
			return nil, err
		}
		fmt.Println("test ListEncodeService: ", string(body))
		return body, nil
	}

	resp := response2.NewResp()

	r.ParseForm()
	offset := r.Form.Get("offset")
	limit := r.Form.Get("limit")
	types := r.Form.Get("type")
	status := r.Form.Get("status")
	serverID := r.Form.Get("server_id")

	body, err := f(offset, limit, types, status, serverID)
	if err != nil {
		log.Println("ListEncodeService err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
