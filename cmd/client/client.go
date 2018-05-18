// Super cheap-ass client for GRPC service.
package main

import (
	"github.com/ianfoo/grpc-math-service/pb"
	"google.golang.org/grpc"

	"context"
	"flag"
	"fmt"
	"os"
)

const defaultAddr = ":3000"

var (
	command = flag.String("operation", "add", "arithmetic to perform: add/subtract/multiply/divide")
	a       = flag.Int64("a", 0, "first input to operation")
	b       = flag.Int64("b", 0, "second input to operation")
)

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = defaultAddr
	}

	flag.Parse()

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error establishing connection: %v\n", err)
		os.Exit(1)
	}

	client := pb.NewMatherClient(conn)
	switch *command {
	case "add":
		resp, err := client.Add(
			context.Background(),
			&pb.AddRequest{A: *a, B: *b})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error adding: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("sum of %d and %d is %d\n", *a, *b, resp.Sum)
	case "subtract":
		resp, err := client.Subtract(
			context.Background(),
			&pb.SubtractRequest{A: *a, B: *b})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error subtracting: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("difference between %d and %d is %d\n", *a, *b, resp.Difference)
	case "multiply":
		resp, err := client.Multiply(
			context.Background(),
			&pb.MultiplyRequest{A: *a, B: *b})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error multiplying: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("product of %d and %d is %d\n", *a, *b, resp.Product)
	case "divide":
		resp, err := client.Divide(
			context.Background(),
			&pb.DivideRequest{Dividend: *a, Divisor: *b})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error dividing: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%d divided by %d is %d, with remainder %d\n",
			*a, *b, resp.Quotient, resp.Remainder)
	default:
		fmt.Fprintf(os.Stderr, "%q is not a supported operation\n", *command)
		os.Exit(1)
	}
}
