package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/udaykirankavaturu/api-communication-styles-demo/grpc-demo/grpc_demo_go/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func main() {
    // --- Protobuf vs JSON size demo ---
    person := &pb.Person{
        Name:  "Alice",
        Id:    123,
        Email: "alice@example.com",
        Tags:  []string{"engineer", "manager", "mentor"},
        Phones: []*pb.PhoneNumber{
            {Number: "123-456-7890", Type: "mobile"},
            {Number: "987-654-3210", Type: "work"},
        },
    }
    protoBytes, err := proto.Marshal(person)
    if err != nil {
        log.Fatalf("failed to marshal proto: %v", err)
    }
    personMap := map[string]interface{}{
        "name":   person.Name,
        "id":     person.Id,
        "email":  person.Email,
        "tags":   person.Tags,
        "phones": []map[string]interface{}{
            {"number": person.Phones[0].Number, "type": person.Phones[0].Type},
            {"number": person.Phones[1].Number, "type": person.Phones[1].Type},
        },
    }
    jsonBytes, err := json.Marshal(personMap)
    if err != nil {
        log.Fatalf("failed to marshal json: %v", err)
    }
    fmt.Printf("Protobuf size: %d bytes\n", len(protoBytes))
    fmt.Printf("JSON size:     %d bytes\n", len(jsonBytes))
    fmt.Printf("\nProtobuf (raw): %v ...\n", protoBytes[:min(50, len(protoBytes))])
    fmt.Printf("JSON (raw): %s ...\n\n", string(jsonBytes[:min(100, len(jsonBytes))]))

    // Write protobuf and JSON to files
    err = os.WriteFile("person.bin", protoBytes, 0644)
    if err != nil {
        log.Fatalf("failed to write protobuf file: %v", err)
    }
    err = os.WriteFile("person.json", jsonBytes, 0644)
    if err != nil {
        log.Fatalf("failed to write JSON file: %v", err)
    }
    fmt.Println("Wrote person.bin (protobuf) and person.json (JSON)")

    // --- Existing gRPC client demo ---
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

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
