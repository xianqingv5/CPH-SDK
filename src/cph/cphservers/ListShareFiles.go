package cphservers

import (
	"fmt"
	"httphelper"
	"net/http"
)

func ListShareFiles(w http.ResponseWriter, r *http.Request)  {
	path := r.Form.Get("path")
	serverIDs := r.Form.Get("server_ids")
	offset := r.Form.Get("offset")
	limit := r.Form.Get("limit")

	if len(path) == 0 || len(serverIDs) == 0 {
		return
	}

	uri := fmt.Sprintf("https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/servers/share-files?offset=%s&limit=%s&path=%s&server_ids=%s", offset, limit, path, serverIDs)
	body, err := httphelper.HttpGet(uri)
	if err != nil {
		return
	}

	WriteTo(w, body)
}