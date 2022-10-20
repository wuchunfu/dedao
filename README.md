# dedao

> 🦉 《得到》 APP web端接口，可查看已购买的课程，听书书架，电子书架，锦囊，推荐话题等

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/yann0917/dedao)

## 特别声明

仅供个人学习使用，请尊重版权，内容版权均为得到所有，请勿传播内容！！！

仅供个人学习使用，请尊重版权，内容版权均为得到所有，请勿传播内容！！！

仅供个人学习使用，请尊重版权，内容版权均为得到所有，请勿传播内容！！！

## 目录结构

```text
.
├── LICENSE
├── README.md
├── app
│   ├── article.go
│   ├── base.go
│   ├── course.go
│   ├── download.go
│   ├── ebook.go
│   ├── login.go
│   └── topic.go
├── config
│   ├── config.go
│   └── dedao.go
├── downloader
│   ├── downloader.go
│   └── types.go
├── go.mod
├── go.sum
├── request
│   ├── download.go
│   └── http.go
├── services // 请求
│   ├── article.go
│   ├── chapter.go
│   ├── course.go
│   ├── course_category.go
│   ├── ebook.go
│   ├── login.go
│   ├── media.go
│   ├── requester.go
│   ├── service.go
│   ├── service_test.go
│   ├── topic.go
│   └── user.go
└── utils
    ├── chromedp.go
    ├── chromedp_test.go
    ├── crypt.go
    ├── ffmpeg.go
    ├── html2epub.go
    ├── json.go
    ├── pool.go
    ├── qrcode.go
    ├── svg2html.go
    ├── utils.go
    └── utils_test.go
```

## 安装

使用如下命令安装：

`go get -u github.com/yann0917/dedao`

## 使用方法

## Stargazers over time

[![Stargazers over time](https://starchart.cc/yann0917/dedao.svg)](https://starchart.cc/yann0917/dedao)

## License

[MIT](./LICENSE) © yann0917

## Support

[![jetbrains](https://s1.ax1x.com/2020/03/26/G9uQoR.png)](https://www.jetbrains.com/?from=dedao)

---
