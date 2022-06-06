package authtoken

import (
	"bytes"
	// "io/ioutil"
	"net/http"
)

// Authtoken 获取token
func Authtoken() string {
	url := "https://iam.ru-northwest-2.myhuaweicloud.com/v3/auth/tokens?nocatalog=true"

	// json序列化
	post := `{
		"auth": {
			"identity": {
				"methods": [
					"password"
				],
				"password": {
					"user": {
						"domain": {
							"name": "diangao-2"
						},
						"name": "lancelot.tian",
						"password": "Yeahmobissp1!"
					}
				}
			},
			"scope": {
				"project": {
					"name": "ru-northwest-2"
				}
			}
		}
	}`

	var jsonStr = []byte(post)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println("status", resp.Status)
	// fmt.Println("response:", resp.Header["X-Subject-Token"])
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
	// fmt.Println(resp.Header["X-Subject-Token"][0])

	return resp.Header["X-Subject-Token"][0]
}
