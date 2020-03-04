package main

import (
	"flag"
	"fmt"
	"github.com/euank/go-kmsg-parser/kmsgparser"
	"oom_notify/utils/notify"
	"oom_notify/utils/oom_match"
	"os"
)

var dingdingtoken string
var result_msg string

func main() {
	hostname, err_get_host := os.Hostname()
	if err_get_host != nil {
		fmt.Println("get hostname error...")
		os.Exit(2)
	}
	// 设置命令行参数
	flag.StringVar(&dingdingtoken, "token", "", "dingding token")
	flag.Parse()

	if len(dingdingtoken) <= 0 {
		fmt.Println("dingdingtoken is too short...")
		os.Exit(3)
	}
	fmt.Println("oom notify start...")

	parser, err := kmsgparser.NewParser()

	if err != nil {
		fmt.Printf("unable to create parser: %v", err)
		os.Exit(4)
	}
	defer parser.Close()
	parse_err := parser.SeekEnd()
	if parse_err != nil {
		fmt.Printf("could not tail: %v", err)
		os.Exit(5)
	}
	kmsg := parser.Parse()

	for msg := range kmsg {
		oom_result := oom_match.OomAnalyse(msg.Message)
		if oom_result.IsOom {
			if oom_result.IsCgroup {
				result_msg = fmt.Sprintf("%s oom happend. type: Cgroup", hostname)
			} else {
				result_msg = fmt.Sprintf("%s oom happend. type: System", hostname)
			}
			notify.Dingding(dingdingtoken, result_msg)
		}

	}
}
