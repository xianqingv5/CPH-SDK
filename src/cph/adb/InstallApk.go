package adb

type PostBody struct {
	Command string `json:"command"`
	Content string `json:"content"`
	ServerIds string `json:"server_ids"`
	PhoneIds string `json:"phone_ids"`
}