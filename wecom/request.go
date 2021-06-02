package wecom

import (
	"bytes"
	"log"
	"net/http"
)

var Gateway = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=038df1a7-4559-4d03-8c6a-cca7c666c3cd"

func sendToWeCom(content []byte) {
	buf := bytes.NewBuffer(content)
	res, err := http.Post(Gateway, "applocation/json", buf)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer res.Body.Close()
	println(res, err)
}
