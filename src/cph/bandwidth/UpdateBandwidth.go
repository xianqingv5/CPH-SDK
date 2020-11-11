package bandwidth

import (
	"encoding/json"
	"fmt"
	"httphelper"
	"net/http"
	"strconv"
)

type bws struct {
	BandWidthSize int `json:"band_width_size"`
}

func UpdateBandwidth(w http.ResponseWriter, r *http.Request) {
	bwID := r.Form.Get("band_width_id")
	if len(bwID) == 0 {
		return
	}

	bwSize := r.Form.Get("band_width_size")
	if len(bwSize) == 0 {
		return
	}
	info, _ := strconv.Atoi(bwSize)
	data := bws{BandWidthSize: info}
	mydata, _ := json.Marshal(data)

	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/bandwidths/" + bwID

	body, err := httphelper.HttpPut(uri, mydata)
	if err != nil {
		return
	}
	fmt.Println("test QueryBandwidth: ", string(body))

	WriteTo(w, body)
}
