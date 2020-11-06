package main

import (
	"cph/cphservers"
	"net/http"
)

func main() {
	http.HandleFunc("/test", cphservers.ListCloudPhoneServers)

	http.ListenAndServe("0.0.0.0:11111", nil)
}
