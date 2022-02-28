## KeyValue

KeyValue is an app to store any key and value you need. It exposes four endpoints:

- `GET /vault` (get a list of all keys and values in the database)
- `GET /vault/{key}` (get a specific key)
- `POST /vault` (post a new key and value; if the key already exists, it updates the value)

```json
{ 
    "key": "test",
    "value": "testsecret"
}
```

- `DELETE /vault/{key}` (delete a specific entry using the key)

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

## About the code

This was the first time that I wrote a code in Go. I first decided to do the test in JavaScript, but since the role is to work with Go, I thought it would be a good way to practice and a great challenge to try and learn Go. After reading the documentation and practicing with the examples on the websites Go By Example, Go Web By Example, and a few tutorials on YouTube, I've decided to use the Gorilla Mux package as the router and the Gorm ORM. 