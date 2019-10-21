package main

import (
	"flag"
	"fmt"
	"github.com/hpcloud/tail"
	"oom_notify/utils/notify"
	"oom_notify/utils/oom_match"
	"os"
)

var dingdingtoken string
var logfile string
var msg string

func main() {
	hostname, err_get_host := os.Hostname()
	if err_get_host != nil {
		fmt.Println("get hostname error...")
		os.Exit(1)
	}
	// 设置命令行参数
	flag.StringVar(&dingdingtoken, "token", "", "dingding token")
	flag.StringVar(&logfile, "logfile", "/var/log/syslog", "system log file")
	flag.Parse()

	if len(dingdingtoken) <= 0 {
		fmt.Println("dingdingtoken is too short...")
		os.Exit(1)
	}
	fmt.Println("oom notify start...")
	t, err := tail.TailFile(logfile, tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		fmt.Printf("oom notify start failed. reason: %s", err)
		os.Exit(1)
	}
	for line := range t.Lines {
		oom_result := oom_match.OomAnalyse(line.Text)
		if oom_result.IsOom {
			if oom_result.IsCgroup {
				msg = fmt.Sprintf("%s oom happend. type: Cgroup", hostname)
			} else {
				msg = fmt.Sprintf("%s oom happend. type: System", hostname)
			}
			notify.Dingding(dingdingtoken, msg)
		}

	}
}
