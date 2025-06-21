# gRPC Demo in Go

This is a simple gRPC demo with a server and client written in Go. It demonstrates basic unary communication.

---

## Prerequisites

- Go 1.16+
- `protoc` (Protocol Buffers compiler)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

Install plugins if you haven't:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Make sure `$GOPATH/bin` is in your `PATH`.

---

## Setup & Run

### 1. Clone and navigate into the project

```bash
cd grpc_demo_go
```

### 2. Initialize Go module

```bash
go mod tidy
```

### 3. Generate gRPC code from .proto file

```bash
protoc --go_out=. --go-grpc_out=. proto/helloworld.proto
```

### 4. Run the server

```bash
go run server/main.go
```

### 5. Run the client (in a separate terminal)

```bash
go run client/main.go Uday
```

Expected output:

```
Greeting: Hello Uday
```

---

## File Structure

```
grpc_demo_go/
├── client/
│   └── main.go
├── proto/
│   └── helloworld.proto
├── server/
│   └── main.go
├── go.mod
├── go.sum
└── README.md
```
