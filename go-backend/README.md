# Tong quan

Du an back end go:

+ trien khai theo mo hinh MVC
+ trong tam la service -> La lop ket noi logic nghiep vu va DB
+ Dung GIN lam REST API

# Project Struct

project/
├── Makefile                 # Script build, run, test, deploy
├── configs/                 # File cấu hình (yaml, json, env) Dung Viper de khoi tao
├── docs/                    # Tài liệu dự án, API doc
├── global/                  # Các config, constants, logger, middleware chung toàn dự án
│
├── internal/                # Code core/private của dự án
│   ├── dao/                 # Data access object, truy cập DB
│   ├── middleware/          # Middleware cho API (auth, logging,…)
│   ├── model/               # Struct model, entity
│   ├── routers/             # Route/handler setup, controller
│   └── service/             # Business logic
│
├── pkg/                     # Package reusable, có thể dùng ngoài module
│
├── storage/                 # File storage, local db, cache, upload, static assets
│
├── scripts/                 # Scripts helper: build, migration, deploy
├── third_party/             # Thư viện ngoài (non-go mod) hoặc các tool add-on
│
├── go.mod                   # Go module
├── .gitignore
└── LICENSE
