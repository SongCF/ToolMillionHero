# ToolMillionHero
百万英雄, 冲顶大会, 知识超人, 答题助手, 自动决策答案

# CHANGELOG
- 2018-01-10：添加图片裁切（减少评论区的文字干扰，减少网络IO）
- 2018-01-11：手机上性能不必电脑高，将执行文件转移到电脑上
- TODO：将页面内容请求改为并发，降低总请求时间
- TODO：自动点击屏幕答题
- TODO：自动识别下一道题，一次运行答所有题 


---

# 原理
### Android
1. 截屏，裁切出图片的题目范围
2. 获取图片内容，然后通过文字识别api（目前使用的百度ocr）获取图片文字
3. 使用百度搜索，并统计搜索得到结果数量
4. 输出频率最高的答案
5. TODO:通过ocr api返回信息识别答案位置，自动点击
6. TODO:自动检测下一道题是否开始

### iOS
1.TODO

---

#使用
### Android 设备
0. 申请一个百度开发账号，注册中文识别api服务（每天免费500次），获取`app_id` `app_key` `secret_key`，并在`baidu/conf.go`中替换它们
1. 电脑一台，编译：`go build android/main.go`
2. Android 手机一台，电脑上安装 ADB，连接上电脑后开启 USB 调试模式
3. 进入百万英雄答题界面执行`run.cmd`，等待输出答案

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
- 个人微信：s821416394
- 百万英雄（交流群）: 522564417

