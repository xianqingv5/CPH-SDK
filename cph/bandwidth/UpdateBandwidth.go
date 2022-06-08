package bandwidth

import (
	"CPH-SDK/conf"
	"CPH-SDK/httphelper2"
	"CPH-SDK/response2"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type bws struct {
	BandWidthSize int `json:"band_width_size"`
}

func UpdateBandwidth(w http.ResponseWriter, r *http.Request) {
	resp := response2.NewResp()

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
	uri := fmt.Sprintf("https://%s/v1/%s/cloud-phone/bandwidths/%s", conf.Config.Huawei.Endpoint, conf.Config.Huawei.ProjectId, bwID)
	// uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/bandwidths/" + bwID

	body, err := httphelper2.HttpPut(uri, mydata)
	if err != nil {
		log.Println("UpdateBandwidth err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
