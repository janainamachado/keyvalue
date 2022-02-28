## KeyValue

KeyValue is an app to store any key and value you need. It exposes four endpoints:

- `GET /vault`
- `GET /vault/{key}`
- `POST /vault`
- `DELETE /vault/{key}`

## Stack

- MySQL
- Go

## Dependencies

- "github.com/go-sql-driver/mysql"
- "github.com/gorilla/mux"
- "github.com/jinzhu/gorm"

## Quick Start

```bash
go run main.go
```
