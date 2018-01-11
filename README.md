# ToolMillionHero
百万英雄, 冲顶大会, 知识超人, 答题助手, 自动决策答案

# TODO
1. 添加百度搜索请求功能


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
1. 编译：`CGO_ENABLED=0 GOARCH=arm GOOS=linux go build -o millionhero android/main.go`
2. Android 手机一台，电脑上安装 ADB，连接上电脑后开启 USB 调试模式
3. 将编译好的文件 Push 到手机上 `adb push ./millionhero /data/local/tmp/`
4. 准备 `adb shell && cd /data/local/tmp/ && chmod 775 ./millionhero`
5. 每当题目出现时执行`./millionhero`等待输出答案


### iOS 设备
1.TODO