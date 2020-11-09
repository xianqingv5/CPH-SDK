package codingservice

import (
	"encoding/json"
	"fmt"
	"httphelper"
	"net/http"
	"strings"
)

type esIDs struct {
	EncodeServerIds []string `json:"encode_server_ids"`
}

func RestartEncodeServer(w http.ResponseWriter, r *http.Request) {
	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/encode-servers/batch-restart"

	esids := r.Form.Get("encode_server_ids")
	if len(esids) == 0 {
		return
	}

	postbody := esIDs{EncodeServerIds: strings.Split(esids, ",")}
	data, _ := json.Marshal(postbody)

	body, err := httphelper.HttpPost(uri, data)
	if err != nil {
		return
	}
	fmt.Println("test RestartEncodeServer: ", string(body))

	WriteTo(w, body)
}
