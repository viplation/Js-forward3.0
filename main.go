package main

import (
	"bufio"
	"fmt"
	_ "jsForward/boot"
	_ "jsForward/router"
	"os"
	"strconv"
	"strings"

	"github.com/gogf/gf/frame/g"
)

func main() {
	fmt.Println(`

	___  _______         _______  _______  ______    _     _  _______  ______    ______  
    |   ||       |       |       ||       ||    _ |  | | _ | ||   _   ||    _ |  |      | 
    |   ||  _____| ____  |    ___||   _   ||   | ||  | || || ||  |_|  ||   | ||  |  _    |
    |   || |_____ |____| |   |___ |  | |  ||   |_||_ |       ||       ||   |_||_ | | |   |
 ___|   ||_____  |       |    ___||  |_|  ||    __  ||       ||       ||    __  || |_|   |
|       | _____| |       |   |    |       ||   |  | ||   _   ||   _   ||   |  | ||       |
|_______||_______|       |___|    |_______||___|  |_||__| |__||__| |__||___|  |_||______| 

	Version 3.0 By kio     
 `)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("============================================================================================")
	fmt.Print(">请输入当前网卡的 IP 地址:")
	currentIP, _ := reader.ReadString('\n')
	currentIP = strings.TrimSpace(currentIP) // 去除空白字符

	fmt.Print(">请输入目标站点协议(HTTP/HTTPS):")
	protocol, _ := reader.ReadString('\n')
	protocol = strings.ToLower(strings.TrimSpace(protocol))

	var FORWORD_PORT int
	var urlPrefix string
	if protocol == "http" {
		FORWORD_PORT = 80
		urlPrefix = "http://"
	} else if protocol == "https" {
		FORWORD_PORT = 443
		urlPrefix = "https://"
	} else {
		fmt.Println(">您的协议输入有误")
		return
	}

	for {
		fmt.Print(">请输入要forward到Burp的参数名(输入$end结束):")
		paramName, _ := reader.ReadString('\n')
		paramName = strings.TrimSpace(paramName)

		if paramName == "$end" {
			break
		}

		fmt.Print(">请输入" + paramName + "的数据类型(json/string):")
		dataType, _ := reader.ReadString('\n')
		dataType = strings.TrimSpace(dataType)

		fmt.Print(">请输入请求标识(例如:REQUEST/RESPONSE):")
		requestType, _ := reader.ReadString('\n')
		requestType = strings.ToUpper(strings.TrimSpace(requestType))

		apiEndpoint := ""
		if requestType == "REQUEST" {
			apiEndpoint = "/api/request"
		} else if requestType == "RESPONSE" {
			apiEndpoint = "/api/response"
		} else {
			fmt.Println(">您的请求标识输入有误")
			continue
		}

		basePayload := ""
		if dataType == "json" {
			basePayload = `var xhr = new XMLHttpRequest();xhr.open("post", "` + urlPrefix + currentIP + `:` + strconv.Itoa(FORWORD_PORT) + apiEndpoint + `", false);xhr.send(JSON.stringify(` + paramName + `));` + paramName + `=JSON.parse(xhr.responseText);`
		} else if dataType == "string" {
			basePayload = `var xhr = new XMLHttpRequest();xhr.open("post", "` + urlPrefix + currentIP + `:` + strconv.Itoa(FORWORD_PORT) + apiEndpoint + `", false);xhr.send(` + paramName + `);` + paramName + `=xhr.responseText;`
		} else {
			fmt.Println(">您的数据类型输入有误")
			continue
		}

		fmt.Println("payload生成完毕:\n" + basePayload)
		fmt.Println("============================================================================================")
	}
	g.Server().Run()
}
