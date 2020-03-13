package v1

import (
	"bilibli-ios/util"
	"fmt"
	"os"
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
	Del    bool `json:"del" comment:"是否删除源文件"`
}

// 合成视频
func (b *Bili) FfmpegAudioVideo() {
	if b.Ignore {
		return
	}
	// 检测目录是否已存在
	if _, err := os.Stat(b.SaveDir); err != nil {
		// 生成目录
		_ = os.MkdirAll(b.SaveDir, 0777)
	}
	fmt.Println("----------------------------------------------------------------")
	fmt.Println(fmt.Sprintf("开始生成目录：%s\r\n 文件名字：%s", b.SaveDir, b.SaveName))
	fmt.Println("----------------------------------------------------------------")
	util.ExecCommand1("/c", fmt.Sprintf("ffmpeg  -i %s -i %s -c:v copy -c:a aac -strict experimental %s", b.Video, b.Audio, b.Save))
	if b.Del {
		_ = os.Remove(b.Audio)
		_ = os.Remove(b.Video)
	}
	fmt.Println("----------------------------------------------------------------")
	fmt.Println(fmt.Sprintf("任务已完成：%s\r\n 文件名字：%s", b.SaveDir, b.SaveName))
	fmt.Println("----------------------------------------------------------------")
}
