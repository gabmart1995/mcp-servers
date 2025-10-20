package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mcp-projects-go/models"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const API_URL = "http://localhost:3000/project"

// salva un nuevo projecto
func saveProject(ctx context.Context, ctr *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	data, err := json.Marshal(ctr.Params.Arguments)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/save", API_URL),
		bytes.NewBuffer(data),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	// extraemos el body
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: string(body),
			},
		},
	}

	return result, nil
}

// permite listar los proyectos en la base de datos
func listProjects(ctx context.Context, ctr *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	request, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/list", API_URL),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	// extraemos el body
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: string(body),
			},
		},
	}

	return result, nil
}

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "mcp-projects-go",
		Version: "1.0.0",
	}, nil)

	server.AddTool(
		&mcp.Tool{
			Name:        "save_project",
			Description: "Crea un nuevo proyecto",
			InputSchema: models.InputSchema{
				Type: "object",
				Properties: map[string]models.Properties{
					"name":        {Type: "string", Description: "nombre del proyecto"},
					"description": {Type: "string", Description: "descripción del proyecto"},
					"state":       {Type: "string", Description: "Estado actual del proyecto"},
				},
				Required: []string{"name", "description", "state"},
			},
		},
		saveProject,
	)

	server.AddTool(
		&mcp.Tool{
			Name:        "list_projects",
			Description: "Lista todos los proyectos",
			InputSchema: models.InputSchema{
				Type:       "object",
				Properties: map[string]models.Properties{},
				Required:   []string{},
			},
		},
		listProjects,
	)

	// corremos el server
	log.Fatal(server.Run(context.Background(), &mcp.StdioTransport{}))
}
