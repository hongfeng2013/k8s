package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			w.Header().Set(k, v[0])
		}
	}
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Println("visitor Info：")
	ip, port, _ := net.SplitHostPort(r.RemoteAddr)
	fmt.Println("IP: ", ip)
	fmt.Println("port: ", port)
	fmt.Println("Status code: ", http.StatusOK)
}

func main() {
	http.HandleFunc("/healthz", healthCheck)
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe("127.0.0.1:80", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}

}
