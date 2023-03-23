# PodcastGPT

这是一个通过chatgpt来实现自动识别播客内容并且总结播客的主旨的工具。

## 目标

PodcastGPT的目标是帮助人们提前了解播客讲了什么，以及更好的做总结，也希望能直接给出播客文本的阅览以及下载。

## 愿景

PodcastGPT的愿景是帮助人们更好地利用播客资源，提高学习效率，节省时间。

## 计划

PodcastGPT的开发计划可以分为以下几个阶段：

一期：

1. 实现从指定的rss获取播客内容并且下载播客内容。
2. 实现语音识别内容并转化为文字。
3. 实现调用chatgpt接口做总结分析。

二期：

1. 实现前端，包括网页前端的样式和功能，能完成上述需求。
2. 实现播客文本的阅览以及下载。

## 结构

```
├── cmd
│   └── main.go
├── internal
│   ├── api
│   │   ├── handlers.go
│   │   └── routes.go
│   ├── config
│   │   └── config.go
│   ├── models
│   │   └── podcast.go
│   ├── repositories
│   │   └── podcast_repository.go
│   ├── services
│   │   ├── chatgpt_service.go
│   │   └── podcast_service.go
│   └── utils
│       ├── rss_parser.go
│       └── text_converter.go
├── migrations
│   ├── 202201010000_init.sql
│   └── migrate.go
├── static
│   ├── css
│   └── js
├── templates
│   ├── index.html
│   └── podcast.html
├── vendor
├── .env
├── .gitignore
├── README.md
└── go.mod
```

```
cmd: 包含一个或多个main函数，用于生成可执行文件。
internal: 包含项目的内部代码。
api: 包含处理HTTP请求的控制器和路由。
config: 包含应用程序的配置参数。
models: 包含与项目中使用的模型相关的结构体。
repositories: 包含与数据库访问相关的代码。
services: 包含业务逻辑层的代码。
utils: 包含项目中的工具类代码。
migrations: 数据库的迁移文件和迁移脚本等。
static: 包含应用程序的静态文件，如CSS、JavaScript、图像等。
templates: 包含应用程序的HTML模板文件。
vendor: 第三方依赖包。
.env: 应用程序的环境变量，如数据库连接字符串等。
.gitignore: git仓库忽略文件列表。
README.md: 项目说明文件。
go.mod: Golang项目的模块文件。
```

## 使用

你可以将这个项目克隆到本地，然后运行app.py文件，即可启动项目。如果你有任何问题，请随时问我。

##  贡献

如果你想为这个项目做出贡献，请fork这个项目，然后提交pull request。如果你有任何问题，请随时问我。

如果你有任何问题，请随时问我。
