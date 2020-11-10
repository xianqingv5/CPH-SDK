package bandwidth

import (
	"fmt"
	"httphelper"
	"net/http"
)

func UpdateBandwidth(w http.ResponseWriter, r *http.Request) {
	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/bandwidths"

	body, err := httphelper.HttpGet(uri)
	if err != nil {
		return
	}
	fmt.Println("test UpdateBandwidth: ", string(body))

	WriteTo(w, body)
}
