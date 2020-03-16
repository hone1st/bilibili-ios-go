package v1

import (
	"bilibli-ios/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var mp = map[string]interface{}{}
var config = map[string]string{}
var save bool
var err error

// v1版本的开始方法
func Main() {
	config = util.InitConfig()
	save, err = strconv.ParseBool(config["save"])
	if err != nil {
		log.Fatal("save配置转换失败！需要填写字符串false或者true")
	}
	_ = filepath.Walk(util.DealPath(config["video_path"], false), SetInput)
	_ = filepath.Walk(util.DealPath(config["name_path"], false), SetOut)

	// 开始处理
	do := func(i interface{}) {
		v := i.(*Bili)
		v.FfmpegAudioVideo()
	}
	util.GoFunc(14, do, mp)

	fmt.Println("请输入任意键退出：")
	d := "1"
	_, _ = fmt.Scanln(&d)
}

// 获取所有的源文件地址
func SetInput(path string, info os.FileInfo, err error) error {
	if !info.IsDir() && strings.Contains(info.Name(), ".danmaku") {
		uuid := strings.Replace(info.Name(), ".danmaku", "", -1)
		dir := filepath.Dir(path)
		dVideo := util.DealPath(dir+"/0.mp4", false)
		dAudio := util.DealPath(dir+"/0.wav", false)
		// 修改视频名字
		if _, err := os.Stat(dVideo); err != nil {
			_ = os.Rename(util.DealPath(dir+"/0.section", false), dVideo)
		}
		// 修改音频名字
		if _, err := os.Stat(dAudio); err != nil {
			_ = os.Rename(util.DealPath(dir+"/1.section", false), dAudio)
		}

		b := &Bili{
			Video:    dVideo,
			Audio:    dAudio,
			Save:     "",
			SaveDir:  "",
			SaveName: "",
			Ignore:   false,
			Del:      false,
		}

		if !save {
			b.Del = true
		}

		mp[uuid] = b

	}

	return nil
}

// 设置Input
func SetOut(path string, info os.FileInfo, err error) error {
	// 找到这个配置的文件
	if !info.IsDir() && strings.Contains(info.Name(), ".bilitask") {
		uuid := strings.Replace(info.Name(), ".bilitask", "", -1)
		// 如果不存在这个Uuid就过滤
		if v, ok := mp[uuid]; ok {
			v := v.(*Bili)
			// 获取这个文件的内容
			b, _ := ioutil.ReadFile(path)
			temp := map[string]string{}
			_ = json.Unmarshal(b, &temp)
			// 再获取temp中的argv
			var argv map[string]string
			_ = json.Unmarshal([]byte(temp["argv"]), &argv)

			v.SaveDir = util.DealPath(config["dest"]+"/"+argv["avname"], true)
			v.SaveName = argv["title"] + ".mp4"
			v.Save = util.DealPath(v.SaveDir+"/"+v.SaveName, true)
			// 如果已经存在就过滤
			if _, err := os.Stat(v.Save); err == nil {
				v.Ignore = true
			}
		}
	}
	return nil
}
