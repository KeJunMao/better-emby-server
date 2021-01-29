package utils

import (
	"flag"
	"fmt"
)

func BesUsage() {
	fmt.Print(`bes - better emby server
并发 http/icmp 请求获取公益服地址请求耗时
`)
	flag.PrintDefaults()
}
