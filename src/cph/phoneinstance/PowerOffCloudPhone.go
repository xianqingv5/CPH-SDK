package phoneinstance

import (
	"encoding/json"
	"fmt"
	"global"
	"httphelper"
	"log"
	"net/http"
	"response"
)

type pocpBody struct {
	PhoneIDs []string `json:"phone_ids"`
}

func PowerOffCloudPhone(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResp()

	var projectId string // 必填，项目ID
	var pocp pocpBody
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		resp.BadReq(w)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&pocp)
	if err != nil {
		resp.BadReq(w)
		return
	}

	if len(pocp.PhoneIDs) == 0 {
		resp.BadReq(w)
		return
	}

	data, _ := json.Marshal(pocp)
	uri := fmt.Sprintf("%s/%s/cloud-phone/batch-stop", global.BaseUrl, projectId)
	body, err := httphelper.HttpPost(uri, data)
	if err != nil {
		log.Println("PowerOffCloudPhone err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
