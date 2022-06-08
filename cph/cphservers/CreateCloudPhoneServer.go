package cphservers

import (
	"CPH-SDK/conf"
	"CPH-SDK/httphelper2"
	"CPH-SDK/response2"
	"encoding/json"
	"fmt"
	"log"
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

// 购买系统定义网络云手机服务器
// todo test
func CreateCloudPhoneServer(w http.ResponseWriter, r *http.Request) {
	resp := response2.NewResp()

	if r.Method != "POST" {
		resp.BadReqMethod(w)
		return
	}

	var cp CloudPhone
	if err := json.NewDecoder(r.Body).Decode(&cp); err != nil {
		resp.BadReq(w)
		return
	}

	if len(cp.ServerName) == 0 || len(cp.ServerModelName) == 0 || len(cp.PhoneModelName) == 0 || len(cp.ImageID) == 0 ||
		cp.Count == 0 || cp.BandWidth == nil || cp.ExtendParam == nil {
		resp.BadReq(w)
		return
	}
	uri := fmt.Sprintf("https://%s/v1/%s/cloud-phone/phones", conf.Config.Huawei.Endpoint, conf.Config.Huawei.ProjectId)
	// uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/phones"
	data, _ := json.Marshal(cp)

	body, err := httphelper2.HttpPost(uri, data)
	if err != nil {
		log.Println("CreateCloudPhoneServer err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
