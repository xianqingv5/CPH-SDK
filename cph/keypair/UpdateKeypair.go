package keypair

import (
	"CPH-SDK/conf"
	"CPH-SDK/httphelper2"
	"CPH-SDK/response2"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ukpBody struct {
	Servers []struct {
		KeypairName string `json:"keypair_name"`
		ServerID    string `json:"server_id"`
	} `json:"servers"`
}

func UpdateKeypair(w http.ResponseWriter, r *http.Request) {
	resp := response2.NewResp()

	// var projectId string // 必填，项目ID

	// r.ParseForm()
	// if len(r.Form.Get("projectId")) > 0 {
	// 	projectId = r.Form.Get("projectId")
	// } else {
	// 	resp.BadReq(w)
	// 	return
	// }
	var ukp ukpBody
	err := json.NewDecoder(r.Body).Decode(&ukp)
	if err != nil {
		resp.BadReq(w)
		return
	}

	if len(ukp.Servers) == 0 {
		resp.BadReq(w)
		return
	}
	mydata, _ := json.Marshal(ukp)
	uri := fmt.Sprintf("https://%s/v1/%s/cloud-phone/servers/open-access", conf.Config.Huawei.Endpoint, conf.Config.Huawei.ProjectId)
	// uri := fmt.Sprintf("%s/%s/cloud-phone/servers/open-access", global.BaseUrl, projectId)
	body, err := httphelper2.HttpPut(uri, mydata)
	if err != nil {
		log.Println("UpdateKeypair err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
