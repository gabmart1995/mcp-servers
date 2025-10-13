package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mcp-conversor-go/models"
	"net/http"
	"reflect"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Properties struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type InputSchema struct {
	Type       string                `json:"type"`
	Required   []string              `json:"required"`
	Properties map[string]Properties `json:"properties"`
}

const (
	API_URL       = "https://cdn.moneyconvert.net/api/latest.json"
	BASE_CURRENCY = "USD"
)

/* obtiene el valor de la moneda */
func getMoneyValue(currency string) (float64, error) {
	request, err := http.NewRequest("GET", API_URL, nil)

	if err != nil {
		return 0.00, err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return 0.00, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return 0.00, err
	}

	var data models.ExchangeRequest

	// asignamos al body
	if err := json.Unmarshal(body, &data); err != nil {
		return 0.00, err
	}

	reflectValue := reflect.ValueOf(data.Rates)
	fieldValue := reflectValue.FieldByName(strings.ToUpper(currency))

	// check if field exists
	if !fieldValue.IsValid() {
		return 0.00, errors.New("No se encontro la moneda solicitada: " + currency)
	}

	return fieldValue.Float(), nil
}

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "mcp-money-conversor",
		Version: "1.0.0",
	}, nil)

	server.AddTool(
		&mcp.Tool{
			Name:        "valor moneda",
			Description: "Devuelve el valor actual de una moneda que necesites (USD, EUR, etc)",
			InputSchema: InputSchema{
				Type: "object",
				Properties: map[string]Properties{
					"currency": {Type: "string", Description: "Valor de la moneda"},
				},
				Required: []string{"currency"},
			},
		},
		func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var data struct {
				Currency string `json:"currency"`
			}

			if err := json.Unmarshal(request.Params.Arguments, &data); err != nil {
				return nil, err
			}

			value, err := getMoneyValue(data.Currency)

			if err != nil {
				return nil, err
			}

			result := &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{
						Text: fmt.Sprintf(
							"el valor actual de la moneda %s frente al es: %s es %.6f",
							data.Currency,
							BASE_CURRENCY,
							value,
						),
					},
				},
			}

			return result, nil
		},
	)

	// corremos el server
	log.Fatal(server.Run(context.Background(), &mcp.StdioTransport{}))
}
