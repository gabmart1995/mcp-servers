package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Input struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

func TestAdd(t *testing.T) {
	ctx := context.Background()
	input := Input{
		A: 25.00,
		B: 25.00,
	}

	// creamos el cliente mcp
	client := mcp.NewClient(&mcp.Implementation{Name: "test-mcp-server-stdio", Version: "0.0.1"}, nil)

	// ejecutamos el servidor mcp en segundo plano
	transport := &mcp.CommandTransport{Command: exec.Command("go", "run", "main.go")}
	session, err := client.Connect(ctx, transport, nil)

	// verificamos los errores
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	// ejecutamos la herramienta mcp
	params := &mcp.CallToolParams{
		Name:      "add",
		Arguments: input,
	}

	// evaluamos la expresion obtenida
	t.Run(fmt.Sprintf("%2.f+%2.f", input.A, input.B), func(t *testing.T) {
		expected := fmt.Sprintf("%.2f + %.2f = %.2f", input.A, input.B, (input.A + input.B))
		result, err := session.CallTool(ctx, params)

		// verificamos los errores
		if err != nil {
			t.Error(err)
			return
		}

		if result.IsError {
			t.Errorf("error: execute failed")
			return
		}

		var value string

		for _, content := range result.Content {
			value = content.(*mcp.TextContent).Text
		}

		if value != expected {
			t.Errorf("Add (%.2f + %.2f) = %s; expected %s", input.A, input.B, value, expected)
		}
	})
}

func TestMultiply(t *testing.T) {
	ctx := context.Background()
	input := Input{
		A: 25.00,
		B: 5.00,
	}

	// creamos el cliente mcp
	client := mcp.NewClient(&mcp.Implementation{Name: "test-mcp-server-stdio", Version: "0.0.1"}, nil)

	// ejecutamos el servidor mcp en segundo plano
	transport := &mcp.CommandTransport{Command: exec.Command("go", "run", "main.go")}
	session, err := client.Connect(ctx, transport, nil)

	// verificamos los errores
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	// ejecutamos la herramienta mcp
	params := &mcp.CallToolParams{
		Name:      "multiply",
		Arguments: input,
	}

	// evaluamos la expresion obtenida
	t.Run(fmt.Sprintf("%2.f*%2.f", input.A, input.B), func(t *testing.T) {
		expected := fmt.Sprintf("%.2f * %.2f = %.2f", input.A, input.B, (input.A * input.B))
		result, err := session.CallTool(ctx, params)

		// verificamos los errores
		if err != nil {
			t.Error(err)
			return
		}

		if result.IsError {
			t.Errorf("error: execute failed")
			return
		}

		var value string

		for _, content := range result.Content {
			value = content.(*mcp.TextContent).Text
		}

		if value != expected {
			t.Errorf("multiply (%.2f * %.2f) = %s; expected %s", input.A, input.B, value, expected)
		}
	})
}

func TestServerInfo(t *testing.T) {
	ctx := context.Background()

	// creamos el cliente mcp
	client := mcp.NewClient(&mcp.Implementation{Name: "test-mcp-server-stdio", Version: "0.0.1"}, nil)

	// ejecutamos el servidor mcp en segundo plano
	transport := &mcp.CommandTransport{Command: exec.Command("go", "run", "main.go")}
	session, err := client.Connect(ctx, transport, nil)

	// verificamos los errores
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	// ejecutamos la herramienta mcp
	params := &mcp.CallToolParams{
		Name:      "get_server_info",
		Arguments: nil,
	}

	t.Run("get_server_info", func(t *testing.T) {
		expected, err := json.Marshal(ServerInfoSchema{
			ServerName:   "example-stdio-server",
			Version:      "0.0.1",
			Transport:    "stdio",
			Capabilities: []string{"tools"},
			Description:  "Example MCP server using stdio transport (MCP 2025-06-18 specification)",
		})

		// verificamos los errores
		if err != nil {
			t.Error(err)
			return
		}

		result, err := session.CallTool(ctx, params)

		// verificamos los errores
		if err != nil {
			t.Error(err)
			return
		}

		if result.IsError {
			t.Errorf("error: execute failed")
			return
		}

		var value string

		for _, content := range result.Content {
			value = content.(*mcp.TextContent).Text
		}

		if value != string(expected) {
			t.Errorf("Error processing %v; expected %s", value, string(expected))
		}
	})
}

func TestGreeting(t *testing.T) {
	ctx := context.Background()

	// creamos el cliente mcp
	client := mcp.NewClient(&mcp.Implementation{Name: "test-mcp-server-stdio", Version: "0.0.1"}, nil)

	// ejecutamos el servidor mcp en segundo plano
	transport := &mcp.CommandTransport{Command: exec.Command("go", "run", "main.go")}
	session, err := client.Connect(ctx, transport, nil)

	// verificamos los errores
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	// ejecutamos la herramienta mcp
	params := &mcp.CallToolParams{
		Name: "get_greeting",
		Arguments: struct {
			Name string `json:"name"`
		}{Name: "Gabriel"},
	}

	t.Run("get_greeting", func(t *testing.T) {
		expected := fmt.Sprintf(
			"Hello %s! Welcome to the MCP stdio server.",
			"Gabriel",
		)

		result, err := session.CallTool(ctx, params)

		// verificamos los errores
		if err != nil {
			t.Error(err)
			return
		}

		if result.IsError {
			t.Errorf("error: execute failed")
			return
		}

		var value string

		for _, content := range result.Content {
			value = content.(*mcp.TextContent).Text
		}

		if value != expected {
			t.Errorf("Error processing %v; expected %s", value, expected)
		}
	})
}
