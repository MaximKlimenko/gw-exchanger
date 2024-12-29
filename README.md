# GW-Exchanger

**GW-Exchanger** — это сервис для управления курсами валют. Сервис предоставляет курсы валют через gRPC и использует PostgreSQL в качестве базы данных.

Структура проекта
```
gw-exchanger/
├── cmd/
│     └── main.go
├── internal/
│     ├── storages/
│     │   ├── model.go
│     │   └── postgres/
│     │        ├── connector.go
│     │        └── methods.go
│     ├── config/
│     │    ├── config.go
│     │    └── defaults.go
│     └── exchanger/
│          └── exchanger.go
├── go.mod
├── Dockerfile
├── config.env
├── Makefile
└── README.md
```
