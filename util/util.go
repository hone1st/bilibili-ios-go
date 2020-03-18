package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"log"
	"math/rand"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"
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

// 开启多协程  返回map的k和v
// wgs 设定每次执行的协程数量
// do 需要协程中执行的方法 传入map的k和v
// data 数据源
func GoFunc(wgs int, do func(...interface{}), data map[interface{}]interface{}) {
	ch := make(chan []interface{}, len(data))
	wg := sync.WaitGroup{}
	wg.Add(wgs)
	go func() {
		defer wg.Done()
		defer close(ch)
		for k, v := range data {
			ch <- []interface{}{k, v}
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
					do(val...)
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

type randItem struct {
	start decimal.Decimal
	end   decimal.Decimal
}

// 获取随机数 按比例返回
// items 自定义key为要返回的值 v为占的百分比 浮点数
// 返回key的值  随机按比例
func GetRandItem(items map[interface{}]float64) (interface{}, error) {
	rand.Seed(time.Now().UnixNano() + rand.Int63())
	c := decimal.NewFromFloat(float64(rand.Intn(100)))
	at := decimal.NewFromInt(0)
	mp := map[interface{}]*randItem{}
	check := 0.00
	for k, f := range items {
		if f <= 0 {
			continue
		}

		mp[k] = &randItem{
			start: at,
			end:   at.Add(decimal.NewFromFloat(f).Mul(decimal.NewFromInt(100))),
		}
		at = mp[k].end
		check += f
	}

	if check > 1.00 {
		return nil, errors.New("所有的占比不能大于1")
	}
	for k, f := range mp {
		if f.start.IntPart() <= c.IntPart() && f.end.IntPart() > c.IntPart() {
			return k, nil
		}
	}
	return nil, nil
}
