package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Server status test handler, only response an JSON result.
func TestStatus(w http.ResponseWriter, req *http.Request) {
	ipAddr := req.RemoteAddr
	if len(ipAddr) == 0 {
		ipAddr = req.Header.Get("X-Forwarded-For")
	}
	res := make(map[string]interface{})
	res["ipAddr"] = ipAddr
	log.Println("ip : ", ipAddr)
	res["time"] = time.Now().String()
	for k, v := range req.Header {
		res[k] = v
		log.Println(k, " : ", v)
	}
	res["status"] = "OK"
	j, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "{\"status\":\"error\"}")
	} else {
		fmt.Fprintln(w, string(j))
	}
}

// url for mobile test
func MobileUrl(w http.RequestWriter, req *http.Request) {

}
