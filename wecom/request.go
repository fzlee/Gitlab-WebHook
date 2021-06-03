package wecom

import (
	"bytes"
	"fmt"
	"gitlab_webhook/utils"
	"log"
	"net/http"
)

var Gateway = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s"

func sendToWeCom(content []byte) {
	url := fmt.Sprintf(Gateway, *utils.WECOM_TOKEN)

	buf := bytes.NewBuffer(content)
	res, err := http.Post(url, "applocation/json", buf)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer res.Body.Close()
	println(res, err)
}
