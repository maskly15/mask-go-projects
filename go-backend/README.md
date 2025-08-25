
# Project Struct 

project/
├── api/                     # API layer: HTTP handler, router, middleware
│   ├── handler/             # Xử lý request (Controller trong MVC)
│   ├── middleware/          # Middleware (auth, logging, CORS,…)
│   └── router.go            # Khởi tạo router, register route
│
├── cmd/                     # Entry points của app
│   ├── app/                 # Khởi động app (HTTP server, gRPC,…)
│   │   └── main.go
│   └── cli/                 # Command line tools (migration, seed,…)
│       └── main.go
│
├── config/                  # File cấu hình, env
│   └── config.yaml/.env
│
├── internal/                # Code private của dự án
│   ├── model/               # Struct model, ORM, entity
│   ├── repository/          # Truy cập DB
│   ├── service/             # Business logic
│   └── util/                # Hàm tiện ích, helper functions
│
├── migrations/              # Database migration scripts
│
├── pkg/                     # Code public, reusable package (không phụ thuộc internal)
│
├── scripts/                 # Scripts build, deploy, install
│
├── test/                    # Test riêng biệt, có thể unit hoặc integration
│
├── web/                     # Front-end code
│
├── go.mod                   # Go module
├── .gitignore               # Git ignore
└── LICENSE

