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
	fmt.Printf("get request %s to %s \n", req.RemoteAddr, req.Host)
	// response
	rsp.WriteHeader(200)
	if _, err := rsp.Write([]byte("request received")); err != nil {
		// deal with "write" error
	}
}
