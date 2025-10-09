package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

/**
Este modulo permite construir MCP mas personalizados
puedes integrar validaciones con go Playgrond Validator
*/

type Properties struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type InputSchema struct {
	Type       string                `json:"type"`
	Required   []string              `json:"required"`
	Properties map[string]Properties `json:"properties"`
}

type ServerInfoSchema struct {
	ServerName   string   `json:"server_name"`
	Version      string   `json:"version"`
	Transport    string   `json:"transport"`
	Capabilities []string `json:"capabilities"`
	Description  string   `json:"description"`
}

func main() {
	server := mcp.NewServer(
		&mcp.Implementation{Name: "mcp-stdio-server", Version: "v1.0.0"},
		nil,
	)

	server.AddTool(
		&mcp.Tool{
			Name:        "add",
			Description: "Add two numbers",
			InputSchema: InputSchema{
				Type: "object",
				Properties: map[string]Properties{
					"a": {Type: "number", Description: "first number"},
					"b": {Type: "number", Description: "second number"},
				},
				Required: []string{"a", "b"},
			},
		},
		func(ctx context.Context, ctr *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				A float64 `json:"a"`
				B float64 `json:"b"`
			}

			if err := json.Unmarshal(ctr.Params.Arguments, &params); err != nil {
				return nil, fmt.Errorf("invalid input %v", err)
			}

			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{
						Text: fmt.Sprintf("%.2f + %.2f = %.2f", params.A, params.B, (params.A + params.B)),
					},
				},
			}, nil
		},
	)

	server.AddTool(
		&mcp.Tool{
			Name:        "multiply",
			Description: "multiply two numbers",
			InputSchema: InputSchema{
				Type: "object",
				Properties: map[string]Properties{
					"a": {Type: "number", Description: "first number"},
					"b": {Type: "number", Description: "second number"},
				},
				Required: []string{"a", "b"},
			},
		},
		func(ctx context.Context, ctr *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				A float64 `json:"a"`
				B float64 `json:"b"`
			}

			if err := json.Unmarshal(ctr.Params.Arguments, &params); err != nil {
				return nil, fmt.Errorf("invalid input %v", err)
			}

			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{
						Text: fmt.Sprintf("%.2f * %.2f = %.2f", params.A, params.B, (params.A * params.B)),
					},
				},
			}, nil
		},
	)

	server.AddTool(
		&mcp.Tool{
			Name:        "get_server_info",
			Description: "Generate information about this mcp server",
			InputSchema: InputSchema{
				Type:       "object",
				Properties: map[string]Properties{},
				Required:   []string{"a", "b"},
			},
		},
		func(ctx context.Context, ctr *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			data, _ := json.Marshal(ServerInfoSchema{
				ServerName:   "example-stdio-server",
				Version:      "0.0.1",
				Transport:    "stdio",
				Capabilities: []string{"tools"},
				Description:  "Example MCP server using stdio transport (MCP 2025-06-18 specification)",
			})

			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{
						Text: string(data),
					},
				},
			}, nil
		},
	)

	server.AddTool(
		&mcp.Tool{
			Name:        "get_greeting",
			Description: "Generate greeting personalized",
			InputSchema: InputSchema{
				Type: "object",
				Properties: map[string]Properties{
					"name": {Type: "string", Description: "Name of person"},
				},
				Required: []string{"name"},
			},
		},
		func(ctx context.Context, ctr *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				Name string `json:"name"`
			}

			if err := json.Unmarshal(ctr.Params.Arguments, &params); err != nil {
				return nil, fmt.Errorf("invalid input %v", err)
			}

			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{
						Text: fmt.Sprintf(
							"Hello %s! Welcome to the MCP stdio server.",
							params.Name,
						),
					},
				},
			}, nil
		},
	)

	log.Fatal(server.Run(context.Background(), &mcp.StdioTransport{}))
}
