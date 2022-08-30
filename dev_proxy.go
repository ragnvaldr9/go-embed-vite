package vite

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type devProxy struct {
	url string
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func (dp *devProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	r.RequestURI = ""
	var err error

	r.URL, err = url.Parse(dp.url + "/" + r.URL.Path)

	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(r)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
