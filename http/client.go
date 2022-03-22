package http

import "net/http"

func StartClient(url string) (*http.Response, error) {
	rsp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}
