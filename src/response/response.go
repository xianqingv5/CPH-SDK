package response

import (
	"encoding/json"
	"global"
	"net/http"
)

// 数据返回格式
type Resp struct {
	Status int         `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}

func NewResp() *Resp {
	return &Resp{
		Status: global.StatusOK,
		Info:   global.StatusText(global.StatusOK),
	}
}

func (r *Resp) WriteTo(w http.ResponseWriter) {
	b, _ := json.Marshal(r)
	w.Write(b)
}

func (r *Resp) BadReq(w http.ResponseWriter) {
	r.Status = global.StatusBadRequest
	r.Info = global.StatusText(r.Status)
	r.WriteTo(w)
}

func (r *Resp) IntervalServErr(w http.ResponseWriter) {
	r.Status = global.StatusInternalServerError
	r.Info = global.StatusText(r.Status)
	r.WriteTo(w)
}

func (r *Resp) BadReqMethod(w http.ResponseWriter)  {
	r.Status = global.StatusMethodNotAllowed
	r.Info = global.StatusText(r.Status)
	r.WriteTo(w)
}
