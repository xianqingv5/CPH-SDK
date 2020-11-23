package keypair

import (
	"encoding/json"
	"fmt"
	"global"
	"log"
	"net/http"

	"httphelper"
	"response"
)

type ukpBody struct {
	Servers []struct {
		KeypairName string `json:"keypair_name"`
		ServerID    string `json:"server_id"`
	} `json:"servers"`
}

func UpdateKeypair(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResp()

	var projectId string // 必填，项目ID
	var ukp ukpBody
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		resp.BadReq(w)
		return
	}

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

	uri := fmt.Sprintf("%s/%s/cloud-phone/servers/open-access", global.BaseUrl, projectId)
	body, err := httphelper.HttpPut(uri, mydata)
	if err != nil {
		log.Println("UpdateKeypair err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
