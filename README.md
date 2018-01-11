# ToolMillionHero
百万英雄, 冲顶大会, 知识超人, 答题助手, 自动决策答案


# 原理
### Android
1. 截屏，裁切出图片的题目范围
2. 获取图片内容，然后通过文字识别api（目前使用的百度ocr）获取图片文字
3. 使用百度搜索，并统计搜索得到结果数量
4. 输出频率最高的答案

### iOS
1.TODO

---

#使用
### Android 设备
1. 电脑一台，编译：`go build android/main.go`
2. Android 手机一台，电脑上安装 ADB，连接上电脑后开启 USB 调试模式
3. 每当题目出现时执行``等待输出答案run.cmd

```$shell
D:\_gopath\src\github.com\SongCF\ToolMillionHero>run.cmd
2018/01/11 22:17:50 image size: 78492
2018/01/11 22:17:52 getQuestion failed
13  牛顿
18  安培
20  法拉第

D:\_gopath\src\github.com\SongCF\ToolMillionHero>run.cmd
2018/01/11 22:21:40 image size: 89180
8  146
24  145
96  147
```


### iOS 设备
1.TODO


---


### 欢迎交流：
个人微信：s821416394
百万英雄（交流群） ：522564417

