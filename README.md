# goadmin

[![golang](https://img.shields.io/badge/Language-Go-green.svg?style=flat)](https://golang.org)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)

Go开发后台管理系统，将前端资源一起编译，是不是很酷？！

![image](./page_home.png)

![image](./page_users.png)

## Features

- 页面布局：[ElementUI](https://element.eleme.cn/#/zh-CN)
- Go开发库：[yiigo](https://github.com/shenghui0779/yiigo)
- 编译前端资源：[go.rice](https://github.com/GeertJohan/go.rice)

## Requirements

`Go1.11+`

## Usage

- 下载 `goadmin` 并解压
- 执行 `go mod tidy`
- 执行 `sh ent.sh`
- 创建数据库后导入 `demo.sql`
- 在 `cmd` 目录下创建配置文件 `.env` 并配置数据库连接，参考 `.env.example`

#### 编译前端资源

- 安装 `go.rice` 工具，参考 [go.rice](https://github.com/GeertJohan/go.rice)
- 在 `cmd` 目录下执行 `rice embed-go`
- 最后 `go build -o goadmin`

> ⚠️ 注意
> 默认登录账号：admin admin
