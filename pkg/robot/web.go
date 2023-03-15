package robot

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func RunServer(port int, configList []Config) {
	for _, config := range configList {
		robot := NewRobot(config)
		log.Println("new robot:", robot, port)
		http.HandleFunc(config.URL, func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "robot", robot)
			handleRequest(w, r.WithContext(ctx))
		})
	}
	// 启动服务器
	addr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(addr, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 解析请求参数
	// 获取请求方法、URL、协议版本等信息
	method := r.Method
	if method == "POST" {
		// 获取请求体信息
		body, _ := ioutil.ReadAll(r.Body)
		var dingMsg DingMsg
		err := json.Unmarshal(body, &dingMsg)
		if err != nil {
			fmt.Fprintf(w, "error msg")
		}
		url := r.URL.String()

		// 输出访问者的全部信息
		log.Printf("URL: %s\nBody:\n%s\n", url, body)
		robot := r.Context().Value("robot").(*Robot)
		SenderId := dingMsg.SenderID
		err = robot.ChatGPT(dingMsg.Text.Content, SenderId, false)
		if err != nil {
			log.Println("send dingTalk message failed", err)
		}
		fmt.Fprintf(w, "OK")
	}
}
