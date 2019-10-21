package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DingdingResponse struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func Dingding(token string, msg string) {
	client := &http.Client{}
	req := `{"msgtype": "text", "text": {"content":"` + msg + `"}}`
	fmt.Println(req)
	req_new := bytes.NewBuffer([]byte(req))
	request, _ := http.NewRequest("POST", "https://oapi.dingtalk.com/robot/send?access_token="+token, req_new)
	request.Header.Set("Content-type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("消息发送失败: %s\n", err.Error())
	}
	if response.StatusCode == 200 {
		p := &DingdingResponse{}
		body, _ := ioutil.ReadAll(response.Body)
		jsonerr := json.Unmarshal([]byte(body), p)
		if jsonerr != nil {
			fmt.Println("dingding响应解析失败")
			return
		}
		if p.Errcode == 0 {
			fmt.Println("dingding消息发送成功")
		} else {
			fmt.Printf("dingding消息发送失败: %s\n", p.Errmsg)
		}

	} else {
		fmt.Println("报警失败...")
	}
}
