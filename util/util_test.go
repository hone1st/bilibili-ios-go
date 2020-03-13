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

	do := func(i interface{}) {

		fmt.Println(i.(*Bili).Video)
	}

	var mp = map[string]interface{}{}

	mp["sdsad"] = &Bili{
		Video: "34242432",
	}

	GoFunc(5, do, mp)
}
