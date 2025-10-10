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

func TestAdd(t *testing.T) {
	var output Output

	ctx := context.Background()
	input := Input{
		A: 25.00,
		B: 25.00,
	}

	// creamos el cliente mcp
	client := mcp.NewClient(&mcp.Implementation{Name: "test-mcp-calculator", Version: "0.0.1"}, nil)

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
		Name: "add",
		Arguments: map[string]float64{
			"a": input.A,
			"b": input.B,
		},
	}

	// evaluamos la expresion obtenida
	t.Run(fmt.Sprintf("%2.f+%2.f", input.A, input.B), func(t *testing.T) {
		expected := "50.00"
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

		for _, content := range result.Content {
			err = json.Unmarshal([]byte(content.(*mcp.TextContent).Text), &output)

			if err != nil {
				t.Error(err)
			}
		}

		if output.Value != expected {
			t.Errorf("Add (%.2f + %.2f) = %s; expected %s", input.A, input.B, output.Value, expected)
		}
	})
}

func TestSubstract(t *testing.T) {
	var output Output

	ctx := context.Background()
	input := Input{
		A: 25.00,
		B: 10.00,
	}

	// creamos el cliente mcp
	client := mcp.NewClient(&mcp.Implementation{Name: "test-mcp-calculator", Version: "0.0.1"}, nil)

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
		Name: "substract",
		Arguments: map[string]float64{
			"a": input.A,
			"b": input.B,
		},
	}

	// evaluamos la expresion obtenida
	t.Run(fmt.Sprintf("%2.f-%2.f", input.A, input.B), func(t *testing.T) {
		expected := "15.00"
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

		for _, content := range result.Content {
			err = json.Unmarshal([]byte(content.(*mcp.TextContent).Text), &output)

			if err != nil {
				t.Error(err)
				return
			}
		}

		if output.Value != expected {
			t.Errorf("Substract (%.2f + %.2f) = %s; expected %s", input.A, input.B, output.Value, expected)
		}
	})
}

func TestMultiply(t *testing.T) {
	var output Output

	ctx := context.Background()
	input := Input{
		A: 25.00,
		B: 5.00,
	}

	// creamos el cliente mcp
	client := mcp.NewClient(&mcp.Implementation{Name: "test-mcp-calculator", Version: "0.0.1"}, nil)

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
		Name: "multiply",
		Arguments: map[string]float64{
			"a": input.A,
			"b": input.B,
		},
	}

	// evaluamos la expresion obtenida
	t.Run(fmt.Sprintf("%2.f*%2.f", input.A, input.B), func(t *testing.T) {
		expected := "125.00"
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

		for _, content := range result.Content {
			err = json.Unmarshal([]byte(content.(*mcp.TextContent).Text), &output)

			if err != nil {
				t.Error(err)
			}
		}

		if output.Value != expected {
			t.Errorf("Multiply (%.2f * %.2f) = %s; expected %s", input.A, input.B, output.Value, expected)
		}
	})
}

func TestDivide(t *testing.T) {
	var output Output

	ctx := context.Background()
	input := Input{
		A: 25.00,
		B: 5.00,
	}

	// creamos el cliente mcp
	client := mcp.NewClient(&mcp.Implementation{Name: "test-mcp-calculator", Version: "0.0.1"}, nil)

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
		Name: "divide",
		Arguments: map[string]float64{
			"a": input.A,
			"b": input.B,
		},
	}

	// evaluamos la expresion obtenida
	t.Run(fmt.Sprintf("%2.f/%2.f", input.A, input.B), func(t *testing.T) {
		expected := "5.00"
		result, err := session.CallTool(ctx, params)

		// verificamos los errores
		if err != nil {
			t.Error(err)
		}

		if result.IsError {
			t.Errorf("error: execute failed")
			return
		}

		for _, content := range result.Content {
			err = json.Unmarshal([]byte(content.(*mcp.TextContent).Text), &output)

			if err != nil {
				t.Error(err)
				return
			}
		}

		if output.Value != expected {
			t.Errorf("Divide (%.2f / %.2f) = %s; expected %s", input.A, input.B, output.Value, expected)
		}
	})
}
