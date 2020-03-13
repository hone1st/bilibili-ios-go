// c.json

```json
{
  "video_path": "F:\\bilibili\\.Downloads\\zzdownloadtaskmanagertask\\av",
  "name_path": "F:\\bilibili\\.Downloads\\zzdownloadtaskmanagertaskfile",
  "save": "true",
  "dest": "F:\\bilibili\\bili"
}
```
##### video_path ios导出来的bili视频地址
##### name_path iso导出来视频信息地址
##### save 合并后是否保留源文件
##### dest 合拼后移动到的目录地址


·····

##### 必须要有ffmpeg

```path
解压之后：
    系统环境中在path中新增一条  路径是解压的所在目录   （到bin  ex: xxxxx/bin）
```

##### bug
```bug
目前有的问题：
    要想合并，磁盘大小要大于2倍原文件
    
    执行时间较长：
        40G的缓存大概快1个小时合并完成
    
直接执行main需要 可能合并视频导致失败，建议打包之后 管理员权限执行
```

##### 适用范围 ios 越狱  
##### bili ios版本5.54.1
##### 提取bili的缓存文件方法
```
路径：
    /var/mobile/Containers/Data/Application/xxxxxxxx/.Downloads  整个目录导出 
推荐工具：
    沙漏助手
```
##### 目录构造和生成的目录
``` PATH
缓存目录
--  F:\\bilibili\\.Downloads\\zzdownloadtaskmanagertask\\av
--  F:\\bilibili\\.Downloads\\zzdownloadtaskmanagertaskfile

目标目录：
-- F:\\bilibili\\bili
       -- \\集合文件夹1\\具体该集合下的视频.mp4
       -- \\集合文件夹2\\具体该集合下的视频.mp4
```


##### 免责声明
```
    不可用作商业用途，仅提供学习，参考！
```




