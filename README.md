# Project shop-dev

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```

## Kiến trúc truyền thống MVC
### Cấu trúc thư mục MVC
Controller + VO/Service + BO/Repository + entity

```bash
server
├── cmd              # start server
├── common           # common
├── configs          # Cấu hình tham số của biến
├── database         # Cơ sở dữ liệu kết nối
├── go.mod
├── internal
│   ├── controller   # controller
│   ├── vo           # VO (View layer)
│   ├── repository   # Truy cập lớp hiển thị
│   └── entity       # Entity
│   ├── pkg          # Nội bộ thư viện
│   └── service      # Lớp logic nghiệp vụ
│   └── bo           # BO (Business Object)
├── main.go          # start server
├── pkg              # bên ngoài thư viện
├── router           # định tuyến hệ thống
└── test             # kiểm tra thư mục

```bash
project
├── Makefile
├── configs
├── docs
├── global # biến toàn cục tham chiếu đến internal như là mongodb, mysql
├── internal
│   ├── repo
│   ├── middleware
│   ├── model
│   ├── routers
│   └── service
├── pkg
├── storage
├── scripts
└── third_pary #swager, ui

