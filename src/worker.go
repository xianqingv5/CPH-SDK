package main

import (
	"cph/bandwidth"
	"cph/codingservice"
	"cph/cphservers"
	"net/http"
)

func main() {
	http.HandleFunc("/test", bandwidth.QueryBandwidth)
	http.HandleFunc("/cphserver", cphservers.RestartCloudPhoneServer)
	http.HandleFunc("/code", codingservice.ListEncodeService)

	http.ListenAndServe("0.0.0.0:11111", nil)
}
