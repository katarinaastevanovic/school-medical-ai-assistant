package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewProxy(target string) http.Handler {
	targetURL, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(targetURL)
}
