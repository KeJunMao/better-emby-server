package utils

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func GetServerList(filePath string) []string {
	text, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var serverList []string
	for _, api := range strings.Split(string(text), "\n") {
		re := regexp.MustCompile("(https?|http)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]")
		match := re.FindString(api)
		if strings.Contains(api, "http") && len(match) > 0 {
			serverList = append(serverList, match)
		}
	}
	return serverList
}
