package bandwidth

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"httphelper"
	"response"
)

type bws struct {
	BandWidthSize int `json:"band_width_size"`
}

func UpdateBandwidth(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResp()

	r.ParseForm()
	bwID := r.Form.Get("band_width_id")
	if len(bwID) == 0 {
		resp.BadReq(w)
		return
	}

	bwSize := r.Form.Get("band_width_size")
	if len(bwSize) == 0 {
		resp.BadReq(w)
		return
	}
	info, _ := strconv.Atoi(bwSize)
	data := bws{BandWidthSize: info}
	mydata, _ := json.Marshal(data)

	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/bandwidths/" + bwID

	body, err := httphelper.HttpPut(uri, mydata)
	if err != nil {
		log.Println("UpdateBandwidth err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
