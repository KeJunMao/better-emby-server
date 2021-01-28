package main

import (
	"flag"
	"fmt"

	"github.com/KeJunMao/better-emby-server/utils"
)

var filePath = flag.String("f", "", "厂妹发给你的服务器消息另存为文件的路径")
var proxyUrl = flag.String("p", "", "使用指定的代理测试")

func init() {
	flag.Usage = utils.BesUsage
	flag.Parse()
}

func main() {
	if *filePath == "" {
		flag.Usage()
		return
	}
	serverList := utils.GetServerList(*filePath)
	ch := make(chan utils.ResultItem)
	for _, api := range serverList {
		go func(s string) {
			result := utils.HttpGet(utils.HttpGetOptions{
				Api:   s,
				Proxy: *proxyUrl,
			})
			ch <- result
		}(api)
	}
	for r := range ch {

		fmt.Println(fmt.Sprintf("%s %dms %s", r.Api, r.Timer, r.Message))
	}
}
