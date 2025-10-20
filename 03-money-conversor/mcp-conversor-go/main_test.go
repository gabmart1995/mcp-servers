package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestMoneyValue(t *testing.T) {
	ctx := context.Background()

	// creamos el cliente mcp
	client := mcp.NewClient(&mcp.Implementation{Name: "test-mcp-money-convertor", Version: "0.0.1"}, nil)

	// ejecutamos el servidor mcp en segundo plano
	transport := &mcp.CommandTransport{Command: exec.Command("go", "run", "main.go")}
	session, err := client.Connect(ctx, transport, nil)

	// verificamos los errores
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	// evaluamos la expresion obtenida
	t.Run("get_money_value", func(t *testing.T) {
		input := struct {
			Currency string `json:"currency"`
		}{Currency: "EUR"}

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

		// ejecutamos la herramienta mcp
		params := &mcp.CallToolParams{
			Name:      "valor_moneda",
			Arguments: input,
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

		if value != expected {
			t.Errorf("%s; expected %s", value, expected)
		}
	})
}

func TestConvertionMoney(t *testing.T) {
	ctx := context.Background()
	input := struct {
		Origin      string
		Destination string
		Amount      float64
	}{Origin: "EUR", Destination: "USD", Amount: 75.50}

	// creamos el cliente mcp
	client := mcp.NewClient(&mcp.Implementation{Name: "test-mcp-money-convertor", Version: "0.0.1"}, nil)

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
		Name:      "conversor_tipo_cambio",
		Arguments: input,
	}

	// evaluamos la expresion obtenida
	t.Run("get_money_conversion", func(t *testing.T) {
		result, err := getCovertion(input)

		// verificamos los errores
		if err != nil {
			t.Error(err)
			return
		}

		expected := fmt.Sprintf(
			"%.2f %s = %.2f %s (Tasa: %.6f, Moneda Base %s)",
			input.Amount,
			input.Origin,
			result.Value,
			input.Destination,
			result.Rate,
			BASE_CURRENCY,
		)

		response, err := session.CallTool(ctx, params)

		// verificamos los errores
		if err != nil {
			t.Error(err)
			return
		}

		if response.IsError {
			t.Errorf("error: execute failed")
			return
		}

		var value string

		for _, content := range response.Content {
			value = content.(*mcp.TextContent).Text
		}

		// se descubrio que los calculos en float 64
		// no son precisos, no es un problema del lenguaje de programacion
		// sino del procesador asi que esta prueba puede fallar
		if value != expected {
			t.Errorf("%s; expected %s", value, expected)
		}
	})

	t.Run("get_money_conversion_same_currency", func(t *testing.T) {
		input.Origin = "USD"
		input.Destination = "EUR"

		result, err := getCovertion(input)

		// verificamos los errores
		if err != nil {
			t.Error(err)
			return
		}

		expected := fmt.Sprintf(
			"%.2f %s = %.2f %s (Tasa: %.6f, Moneda Base %s)",
			input.Amount,
			input.Origin,
			result.Value,
			input.Destination,
			result.Rate,
			BASE_CURRENCY,
		)

		// actualizamos los valores
		params = &mcp.CallToolParams{
			Name:      "conversor_tipo_cambio",
			Arguments: input,
		}

		response, err := session.CallTool(ctx, params)

		// verificamos los errores
		if err != nil {
			t.Error(err)
			return
		}

		if response.IsError {
			t.Errorf("error: execute failed")
			return
		}

		var value string

		for _, content := range response.Content {
			value = content.(*mcp.TextContent).Text
		}

		// se descubrio que los calculos en float 64
		// no son precisos, no es un problema del lenguaje de programacion
		// sino del procesador asi que esta prueba puede fallar
		if value != expected {
			t.Errorf("%s; expected %s", value, expected)
		}
	})
}
