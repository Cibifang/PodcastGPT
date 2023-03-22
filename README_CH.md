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
project/ - 项目根目录
    app.py - 项目入口文件
    rss_reader.py - 从指定的rss获取播客内容并且下载播客内容
    speech_recognition.py - 实现语音识别内容并转化为文字
    text_to_speech.py - 实现将文字转化为语音
    audio_converter.py - 实现音频格式转换
    api_call.py - 调用chatgpt接口做总结分析
    requirements.txt - 项目依赖
    README.md - 项目说明文档
    templates/ - 网页前端模板目录
        index.html - 网页前端入口文件
    static/ - 网页前端静态文件目录
        css/ - 网页前端样式文件目录
            style.css - 网页前端样式文件
        js/ - 网页前端脚本文件目录
            script.js - 网页前端脚本文件
        data/ - 数据目录
        podcast/ - 播客目录
        audio/ - 音频目录
        text/ - 文本目录
```

## 使用

你可以将这个项目克隆到本地，然后运行app.py文件，即可启动项目。如果你有任何问题，请随时问我。

##  贡献

如果你想为这个项目做出贡献，请fork这个项目，然后提交pull request。如果你有任何问题，请随时问我。

如果你有任何问题，请随时问我。
