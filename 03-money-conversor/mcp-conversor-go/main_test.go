package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestValueMoney(t *testing.T) {
	ctx := context.Background()
	input := struct {
		Currency string `json:"currency"`
	}{Currency: "EUR"}

	// creamos el cliente mcp
	client := mcp.NewClient(&mcp.Implementation{Name: "test-mcp-money-converted", Version: "0.0.1"}, nil)

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
		Name:      "valor moneda",
		Arguments: input,
	}

	// evaluamos la expresion obtenida
	t.Run("get_money_value", func(t *testing.T) {
		data, err := getMoneyValue(input.Currency)

		// verificamos los errores
		if err != nil {
			t.Error(err)
			return
		}

		expected := fmt.Sprintf(
			"el valor actual de la moneda %s frente al es: %s es %.6f",
			input.Currency,
			BASE_CURRENCY,
			data,
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
			t.Errorf("%s; expected %s", value, expected)
		}
	})
}
