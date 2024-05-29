```bash
go run ./cmd/server .

# "google/wire" module
go install github.com/google/wire/cmd/wire@latest
wire ./cmd/server

# "swaggo/swag" module
go install github.com/swaggo/swag/cmd/swag@latest
swag init -pd -g ./cmd/server/main.go -o ./third_party/docs/

# "golang.org/x/tools" module
go install golang.org/x/tools/cmd/goimports@latest
goimports -w .
gofmt -s -w .
```
