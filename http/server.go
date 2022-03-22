package http

import (
	"fmt"
	"net/http"
)

func StartServer(addr string) error {
	http.HandleFunc("/hello", SayHello)
	if err := http.ListenAndServe(addr, nil); err != nil {
		return err
	}
	return nil
}

func SayHello(rsp http.ResponseWriter, req *http.Request) {
	fmt.Println("get request from ", req.Host)
	// response
	rsp.WriteHeader(200)
	if _, err := rsp.Write([]byte("request received")); err != nil {
		// deal with "write" error
	}
}
