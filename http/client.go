package http

import "net/http"

func DoRequest(url string) (*http.Response, error) {
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// request.Header.Set("Host", "specific-host")
	request.Host = "specific-host"
	rsp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}
