# PodcastGPT

This is a tool that uses chatgpt to automatically recognize podcast content and summarize the main points of the podcast.

## Goal

The goal of this project is to help people understand what the podcast is about in advance and to make better summaries. It is also hoped that the text of the podcast can be viewed and downloaded directly.

## Vision

The vision of this project is to help people make better use of podcast resources, improve learning efficiency, and save time.

## Plan

The development plan of this project can be divided into the following stages:

Phase 1:

1. Implement getting podcast content from the specified RSS and downloading podcast content.
2. Implement speech recognition content and convert it to text.
3. Implement calling the chatgpt interface to do summary analysis.

Phase 2:

1. Implement the front-end, including the style and function of the web front-end, which can complete the above requirements.
2. Implement the viewing and downloading of podcast text.

## Project Structure

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

## Usage

You can clone this project to your local machine and then run the app.py file to start the project. If you have any questions, please feel free to ask me.

## Contribution

If you want to contribute to this project, please fork this project and then submit a pull request. If you have any questions, please feel free to ask me.

If you have any questions, please feel free to ask me.
