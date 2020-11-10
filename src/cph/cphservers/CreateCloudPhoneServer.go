package cphservers

import (
	"encoding/json"
	"httphelper"
	"net/http"
)

type CloudPhone struct {
	ServerName      string       `json:"server_name"`
	ServerModelName string       `json:"server_model_name"`
	PhoneModelName  string       `json:"phone_model_name"`
	ImageID         string       `json:"image_id"`
	Count           int          `json:"count"`
	KeypairName     string       `json:"keypair_name"`
	VncEnable       string       `json:"vnc_enable"`
	Ports           []ports      `json:"ports"`
	BandWidth       *bandWidth   `json:"band_width"`
	ExtendParam     *extendParam `json:"extend_param"`
}

type ports struct {
	Name               string `json:"name"`
	ListenPort         int    `json:"listen_port"`
	InternetAccessible string `json:"internet_accessible"`
}

type bandWidth struct {
	BandWidthShareType int `json:"band_width_share_type"`
}

type extendParam struct {
	ChargingMode int `json:"charging_mode"`
	PeriodType   int `json:"period_type"`
	PeriodNum    int `json:"period_num"`
	IsAutoPay    int `json:"is_auto_pay"`
}

func WriteTo(w http.ResponseWriter, data []byte) {
	w.Write(data)
}

func CreateCloudPhoneServer(w http.ResponseWriter, r *http.Request) {
	var cp CloudPhone
	if err := json.NewDecoder(r.Body).Decode(&cp); err != nil {
		return
	}

	if len(cp.ServerName) == 0 || len(cp.ServerModelName) == 0 || len(cp.PhoneModelName) == 0 || len(cp.ImageID) == 0 ||
		cp.Count == 0 || cp.BandWidth == nil || cp.ExtendParam == nil {
		return
	}

	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/phones"
	data, _ := json.Marshal(cp)

	body, err := httphelper.HttpPost(uri, data)
	if err != nil {
		return
	}

	WriteTo(w, body)
}
