package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

var target = os.Getenv("PROXY_TARGET")

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxyURL := url.URL{
			Scheme: "http",
			Host:   target,
		}

		log.Printf("Reverse proxy to: %s", proxyURL.String())
		proxy := httputil.NewSingleHostReverseProxy(&proxyURL)
		proxy.Transport = new(roundTripper)
		proxy.ServeHTTP(w, r)
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Call health")
		w.Write([]byte("OK"))
	})

	http.ListenAndServe(":8081", nil)
}

type roundTripper struct{}

func (r *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Printf("failed to roundtrip: %v", err)
		return nil, err
	}
	oldBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read body: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body := fmt.Sprintf("Called from proxy!!\n%s\n", oldBody)
	resp.Body = ioutil.NopCloser(strings.NewReader(body))
	resp.Header.Del("Content-Length")
	resp.ContentLength = int64(len(body))
	return resp, nil
}
