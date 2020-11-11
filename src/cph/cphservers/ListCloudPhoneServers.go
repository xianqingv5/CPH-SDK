package cphservers

import (
	"net/http"
	"net/url"
	"util"

	"httphelper"
)

func ListCloudPhoneServers(w http.ResponseWriter, r *http.Request) {
	f := func(offset, limit, server_name, server_id string) ([]byte, error) {
		v := url.Values{}
		util.AddurlParam("offset", offset, &v)
		util.AddurlParam("limit", limit, &v)
		util.AddurlParam("server_name", server_name, &v)
		util.AddurlParam("server_id", server_id, &v)

		uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/servers"
		uri = uri + "?" + v.Encode()

		body, err := httphelper.HttpGet(uri)
		return body, err
	}

	r.ParseForm()
	serverName := r.Form.Get("server_name")
	serverID := r.Form.Get("server_id")
	offset := r.Form.Get("offset")
	limit := r.Form.Get("limit")

	body, err := f(offset, limit, serverName, serverID)
	if err != nil {
		return
	}

	WriteTo(w, body)
}
