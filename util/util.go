package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

// 执行cmd命令
func ExecCommand1(cmd ...string) {
	var err error
	if runtime.GOOS == "windows" {
		_, err = exec.Command("cmd", cmd...).Output()
	} else if runtime.GOOS == "linux" {
		_, err = exec.Command(cmd[0], cmd[1:]...).Output()
	}
	if err != nil {
		log.Println("cmd操作出错：", cmd)
	} else {
		fmt.Println("cmd操作完成：", cmd)
	}
}

// 开启多协程
func GoFunc(wgs int, do func(interface{}), data map[string]interface{}) {
	ch := make(chan interface{}, len(data))
	wg := sync.WaitGroup{}
	wg.Add(wgs)
	go func() {
		defer wg.Done()
		defer close(ch)
		for _, v := range data {
			ch <- v
		}
	}()
	for i := 0; i < wgs-1; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case val, ok := <-ch:
					if !ok {
						return
					}
					do(val)
				}
			}
		}()
	}
	wg.Wait()
}

// 处理路径
func DealPath(path string, deal bool) string {
	if runtime.GOOS == "windows" {
		path = strings.Replace(path, "/", "\\", -1)
	} else if runtime.GOOS == "linux" {
		path = strings.Replace(path, "\\", "/", -1)
	}

	if deal {
		path = strings.Replace(path, "&", "、", -1)
		path = strings.Replace(path, " ", "", -1)
	}
	return path
}

// 读取配置
func InitConfig() map[string]string {
	b, _ := ioutil.ReadFile("./c.json")
	config := make(map[string]string, 0)
	_ = json.Unmarshal(b, &config)
	return config
}
