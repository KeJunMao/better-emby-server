package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func HttpGet(options HttpGetOptions) ResultItem {
	result := ResultItem{
		Api:     options.Api,
		Timer:   0,
		Message: "error",
	}
	httpClient := http.DefaultClient
	startT := time.Now()
	if options.Proxy != "" {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(options.Proxy)
		}
		httpTransport := &http.Transport{
			Proxy: proxy,
		}
		httpClient = &http.Client{
			Transport: httpTransport,
		}
	}
	resp, err := httpClient.Get(options.Api)
	if err != nil {
		return result
	}
	defer resp.Body.Close()
	endT := time.Since(startT)
	result.Timer = int(endT.Milliseconds())
	body, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(body), "<title>Emby</title>") {
		result.Message = "ok"
		return result
	}
	result.Message = "not emby"
	return result
}
