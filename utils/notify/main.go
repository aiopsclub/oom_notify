package notify

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Dingding(token string, msg string) {
	client := &http.Client{}
	req := `{"msgtype": "text", "text": {"content":"` + msg + `"}}`
	fmt.Println(req)
	req_new := bytes.NewBuffer([]byte(req))
	request, _ := http.NewRequest("POST", "https://oapi.dingtalk.com/robot/send?access_token="+token, req_new)
	request.Header.Set("Content-type", "application/json")
	response, _ := client.Do(request)
	fmt.Println(response.StatusCode)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	} else {
		fmt.Println("报警失败...")
	}
}
