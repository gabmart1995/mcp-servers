package main

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Input struct {
	A float64 `json:"a" jsonschema:"the value for number 1"`
	B float64 `json:"b" jsonschema:"the value for number 2"`
}

type Output struct {
	Value string `json:"value" jsonschema:"the result for the operation"`
}

func main() {
	server := mcp.NewServer(
		&mcp.Implementation{Name: "mcp-calculator", Version: "0.0.1"},
		nil,
	)

	// add tools
	mcp.AddTool(
		server,
		&mcp.Tool{Name: "add", Description: "add two numbers"},
		func(ctx context.Context, ctr *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, Output, error) {
			value := input.A + input.B
			return nil, Output{Value: strconv.FormatFloat(value, 'f', 2, 64)}, nil
		},
	)

	mcp.AddTool(
		server,
		&mcp.Tool{Name: "substract", Description: "substract two numbers"},
		func(ctx context.Context, ctr *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, Output, error) {
			value := input.A - input.B
			return nil, Output{Value: strconv.FormatFloat(value, 'f', 2, 64)}, nil
		},
	)

	mcp.AddTool(
		server,
		&mcp.Tool{Name: "multiply", Description: "multiply two numbers"},
		func(ctx context.Context, ctr *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, Output, error) {
			value := input.A * input.B
			return nil, Output{Value: strconv.FormatFloat(value, 'f', 2, 64)}, nil
		},
	)

	mcp.AddTool(
		server,
		&mcp.Tool{Name: "divide", Description: "divide two numbers"},
		func(ctx context.Context, ctr *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, Output, error) {
			value := 0.00

			if input.B == 0 {
				return nil, Output{}, errors.New("error: not divide for zero")
			}

			value = input.A / input.B

			return nil, Output{Value: strconv.FormatFloat(value, 'f', 2, 64)}, nil
		},
	)

	// arrancamos el servicio
	log.Fatal(server.Run(context.Background(), &mcp.StdioTransport{}))
}
