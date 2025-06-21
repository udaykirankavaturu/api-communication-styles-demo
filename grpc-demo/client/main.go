package main

import (
    "context"
    "log"
    "os"
    "time"

    "google.golang.org/grpc"
    pb "github.com/udaykirankavaturu/api-communication-styles-demo/grpc-demo/grpc_demo_go/proto"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewGreeterClient(conn)

    name := "world"
    if len(os.Args) > 1 {
        name = os.Args[1]
    }

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.GetMessage())
}
