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

	// se usa el paquete reflect para acceder dinamicamente
	// a las propiedades de la estructura
	reflectValue := reflect.ValueOf(data.Rates)
	fieldValue := reflectValue.FieldByName(strings.ToUpper(currency))

	// check if field exists
	if !fieldValue.IsValid() {
		return 0.00, errors.New("No se encontro la moneda solicitada: " + currency)
	}

	return fieldValue.Float(), nil
}

func getCovertion(params struct {
	Origin      string
	Destination string
	Amount      float64
}) (models.ResultConvert, error) {
	var data models.ExchangeRequest

	result := models.ResultConvert{}
	request, err := http.NewRequest("GET", API_URL, nil)

	if err != nil {
		return result, err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return result, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return result, err
	}

	// asignamos al body
	if err := json.Unmarshal(body, &data); err != nil {
		return result, err
	}

	// obtenemos los campos
	reflectValue := reflect.ValueOf(data.Rates)

	if params.Origin == BASE_CURRENCY {
		fieldValueDestination := reflectValue.FieldByName(strings.ToUpper(params.Destination))

		if !fieldValueDestination.IsValid() {
			return result, errors.New("No se encontro la moneda solicitada: " + params.Destination)
		}

		result.Rate = fieldValueDestination.Float()

	} else {
		fieldValueOrigin := reflectValue.FieldByName(strings.ToUpper(params.Origin))
		fieldValueDestination := reflectValue.FieldByName(strings.ToUpper(params.Destination))

		if !fieldValueOrigin.IsValid() || !fieldValueDestination.IsValid() {
			return result, fmt.Errorf("no se encontraron las monedas solicitadas: %s, %s", params.Origin, params.Destination)
		}

		result.Rate = fieldValueOrigin.Float() / fieldValueDestination.Float()
	}

	result.Value = params.Amount * result.Rate

	return result, nil
}

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "mcp-money-conversor",
		Version: "1.0.0",
	}, nil)

	server.AddTool(
		&mcp.Tool{
			Name:        "valor_moneda",
			Description: "Devuelve el valor actual de una moneda que necesites (USD, EUR, etc)",
			InputSchema: models.InputSchema{
				Type: "object",
				Properties: map[string]models.Properties{
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

	server.AddTool(
		&mcp.Tool{
			Name:        "conversor_tipo_cambio",
			Description: "Devuelve el valor actual de una moneda frente a otra",
			InputSchema: models.InputSchema{
				Type: "object",
				Properties: map[string]models.Properties{
					"origin":      {Type: "string", Description: "Formato inicial"},
					"destination": {Type: "string", Description: "Formato final"},
					"amount":      {Type: "number", Description: "Monto de la conversión"},
				},
				Required: []string{"origin", "destination", "amount"},
			},
		},
		func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				Origin      string  `json:"origin"`
				Destination string  `json:"destination"`
				Amount      float64 `json:"amount"`
			}

			if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
				return nil, err
			}

			result, err := getCovertion(struct {
				Origin      string
				Destination string
				Amount      float64
			}(params))

			if err != nil {
				return nil, err
			}

			response := &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{
						Text: fmt.Sprintf(
							"%.2f %s = %.2f %s (Tasa: %.6f, Moneda Base %s)",
							params.Amount,
							params.Origin,
							result.Value,
							params.Destination,
							result.Rate,
							BASE_CURRENCY,
						),
					},
				},
			}

			return response, nil
		},
	)

	// corremos el server
	log.Fatal(server.Run(context.Background(), &mcp.StdioTransport{}))
}
