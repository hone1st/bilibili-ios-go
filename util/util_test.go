package util

import (
	"fmt"
	"testing"
)

// 非名字都是绝对路径
type Bili struct {
	// 源文件
	Video string `json:"video" comment:"视频文件地址"`
	Audio string `json:"audio" comment:"音频文件地址"`

	// 目标地址
	Save string `json:"save" comment:"合成的文件地址"`
	// 保存视频的目录
	SaveDir  string `json:"save_dir" comment:"保存视频的目录"`
	SaveName string `json:"save_name" comment:"保存文件名字"`
	// 过滤则不合成
	Ignore bool `json:"ignore" comment:"是否过滤"`
}

func TestGoFunc(t *testing.T) {
	do := func(item ...interface{}) {
		fmt.Println(item[1].(int))
		fmt.Println(item[0].(int))
		fmt.Println(item)
	}
	var mp = map[interface{}]interface{}{}

	for i := 0; i < 10; i++ {
		mp[i] = i
	}
	GoFunc(5, do, mp)
}

// 测试返回的数
func TestRand(t *testing.T) {
	do := func(item ...interface{}) {
		mItems := map[interface{}]float64{
			"黄梓健不中奖": 0.9,
			"黄梓健中奖":  0.1,
		}
		i, err := GetRandItem(mItems)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("===", i)
		}
	}
	var mp = map[interface{}]interface{}{}

	for i := 0; i < 10; i++ {
		mp[i] = i
	}
	GoFunc(10, do, mp)

}
