package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type ResultItem struct {
	api     string
	timer   int
	message string
}

func httpGet(api string) ResultItem {
	startT := time.Now()
	resp, err := http.Get(api)
	if err != nil {
		return ResultItem{
			api:     api,
			timer:   9999,
			message: "error",
		}
	}
	defer resp.Body.Close()
	endT := time.Since(startT)
	body, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(body), "<title>Emby</title>") {
		return ResultItem{
			api:     api,
			timer:   int(endT.Milliseconds()),
			message: "ok",
		}
	}
	return ResultItem{
		api:     api,
		timer:   int(endT.Milliseconds()),
		message: "not emby",
	}
}

var filePath = flag.String("f", "", "厂妹发给你的服务器消息另存为文件的路径")

func init() {
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Print(`bes - better emby server

并发 http 请求获取公益服地址请求耗时

`)
	flag.PrintDefaults()
}

func main() {
	if *filePath == "" {
		flag.Usage()
		return
	}
	text, err := ioutil.ReadFile(*filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	var serverList []string
	for _, api := range strings.Split(string(text), "\n") {
		re := regexp.MustCompile("(https?|http)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]")
		match := re.FindString(api)
		if strings.Contains(api, "http") && len(match) > 0 {
			serverList = append(serverList, match)
		}
	}
	ch := make(chan ResultItem)
	for _, api := range serverList {
		go func(s string) {
			result := httpGet(s)
			ch <- result
		}(api)
	}
	for result := range ch {
		fmt.Println(fmt.Sprintf("%s %dms %s", result.api, result.timer, result.message))
	}
}
