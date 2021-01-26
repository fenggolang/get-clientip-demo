package main

import (
	"fmt"
	"net/http"

	"github.com/thinkeridea/go-extend/exnet"
)

func main() {
	fmt.Println("hello")
	http.HandleFunc("/clientip",getClientRealIP)
	http.ListenAndServe(":8080",nil)
}


func getClientRealIP(w http.ResponseWriter,r *http.Request) {
	ip := exnet.ClientPublicIP(r)
	if ip == ""{
		ip = exnet.ClientIP(r)
	}
	w.Write([]byte(ip))
}