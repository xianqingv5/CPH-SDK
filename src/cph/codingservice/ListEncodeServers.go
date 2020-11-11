package codingservice

import (
	"fmt"
	"httphelper"
	"net/http"
	"net/url"

	"util"
)

func WriteTo(w http.ResponseWriter, data []byte) {
	w.Write(data)
}

func ListEncodeService(w http.ResponseWriter, r *http.Request) {
	f := func(offset, limit, types, status, server_id string) ([]byte, error) {
		v := url.Values{}
		util.AddurlParam("offset", offset, &v)
		util.AddurlParam("limit", limit, &v)
		util.AddurlParam("type", types, &v)
		util.AddurlParam("status", status, &v)
		util.AddurlParam("server_id", server_id, &v)

		uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/encode-servers"
		uri = uri + "?" + v.Encode()

		body, err := httphelper.HttpGet(uri)
		if err != nil {
			return nil, err
		}
		fmt.Println("test ListEncodeService: ", string(body))
		return body, nil
	}

	r.ParseForm()
	offset := r.Form.Get("offset")
	limit := r.Form.Get("limit")
	types := r.Form.Get("type")
	status := r.Form.Get("status")
	serverID := r.Form.Get("server_id")

	res, err := f(offset, limit, types, status, serverID)
	if err != nil {
		return
	}

	WriteTo(w, res)
}
