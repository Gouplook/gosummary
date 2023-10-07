package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type FeishuMsgBody struct {
	Card    interface{} `json:"card"`
	MsgType string      `json:"msg_type"`
}

func main() {
	url := "https://open.feishu.cn/open-apis/bot/v2/hook/72b8ac32-00e6-45b1-aae1-0dcb27194c8f"

	card := `{
  "config": {
    "wide_screen_mode": true
  },
  "elements": [
    {
      "tag": "hr"
    },
    {
      "tag": "markdown",
      "content": "\n[明珠 | ASB-5.5KM | rc26]\n(https://www.feishu.cn)\n   + 新增报警：2处（初级2处）\n   + 升级报警：4处（中级2处、高级2处）\n[华电 | AD-5.5KM | rc31]\n(https://www.baidu.com)\n   + 新增报警：2处（初级2处）\n   + 升级报警：4处（中级2处、高级2处）\n\n\n\n",
      "text_align": "left"
    }
  ],
  "header": {
    "template": "red",
    "title": {
      "content": "司图地图质量检查告警系统\n",
      "tag": "plain_text"
    }
  }
}`
	body := FeishuMsgBody{
		MsgType: "interactive",
		Card:    card,
	}

	headers := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
	client := resty.New().SetTimeout(10 * time.Second)
	fResp, err := client.R().SetBody(body).SetHeaders(headers).Post(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fResp.Body()))
	var response Response
	err = json.Unmarshal(fResp.Body(), &response)
	fmt.Println(response)

	fmt.Println("飞书....")
}
