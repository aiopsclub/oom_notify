package oom_match

import (
	"fmt"
	"os"
	"regexp"
)

type oom_info struct {
	IsOom    bool
	IsCgroup bool
}

func OomAnalyse(line string) *oom_info {
	oom_result := &oom_info{}
	oom_re_compile := regexp.MustCompile(`(out|Out)\s+of\s+memory`)
	if oom_re_compile == nil {
		fmt.Println("正则表达式未识别")
		os.Exit(2)
	}

	result := oom_re_compile.FindAllStringSubmatch(line, -1)
	if result != nil {
		oom_result.IsOom = true
	} else {
		oom_result.IsOom = false
	}

	oom_result.IsCgroup = isCgroup(line)

	return oom_result
}

func isCgroup(line string) bool {
	cgroup_re_compile := regexp.MustCompile(`(Cgroup|cgroup)`)
	if cgroup_re_compile == nil {
		fmt.Println("正则表达式未识别")
		os.Exit(3)
	}
	result := cgroup_re_compile.FindAllStringSubmatch(line, -1)
	if result != nil {
		return true
	}
	return false
}
