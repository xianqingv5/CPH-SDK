package cphservers

import (
	"fmt"
	"httphelper"
	"net/http"
	"net/url"
	"util"
)

func ListShareFiles(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	path := r.Form.Get("path")
	serverIDs := r.Form.Get("server_ids")
	offset := r.Form.Get("offset")
	limit := r.Form.Get("limit")

	if len(path) == 0 || len(serverIDs) == 0 {
		return
	}

	v := url.Values{}
	util.AddurlParam("offset", offset, &v)
	util.AddurlParam("limit", limit, &v)

	uri := fmt.Sprintf("https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/servers/share-files?&path=%s&server_ids=%s", path, serverIDs)
	uri = uri + v.Encode()

	body, err := httphelper.HttpGet(uri)
	if err != nil {
		return
	}

	WriteTo(w, body)
}