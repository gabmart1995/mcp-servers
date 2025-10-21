package main

import (
	"context"
	"log"
	"mcp-projects-go/models"
	"os/exec"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestMCPProject(t *testing.T) {
	ctx := context.Background()
	// creamos el cliente mcp
	client := mcp.NewClient(&mcp.Implementation{Name: "test-mcp-projects", Version: "0.0.1"}, nil)

	// ejecutamos el servidor mcp en segundo plano
	transport := &mcp.CommandTransport{Command: exec.Command("go", "run", "main.go")}
	session, err := client.Connect(ctx, transport, nil)

	// verificamos los errores
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	t.Run("list_project_by_id", func(t *testing.T) {
		input := struct {
			Id string `json:"id"`
		}{
			Id: "c903bbc9-8fb4-4a18-9172-ffc16d499d34",
		}

		params := &mcp.CallToolParams{
			Name:      "list_project_id",
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
	})

	t.Run("list_proyects", func(t *testing.T) {
		params := &mcp.CallToolParams{
			Name: "list_project_id",
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
	})

	t.Run("save_project", func(t *testing.T) {
		input := models.Project{
			Name:        "testing project",
			Description: "testing project in react",
			Status:      "progress",
		}

		params := &mcp.CallToolParams{
			Name:      "save_project",
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
	})

	t.Run("update_project", func(t *testing.T) {
		input := models.Project{
			Name:        "testing project",
			Description: "testing project in react",
			Status:      "progress",
			Id:          "04e17f2c-040b-44a4-bf3e-40968779944e",
		}

		params := &mcp.CallToolParams{
			Name:      "save_project",
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
	})

	t.Run("get_image_project", func(t *testing.T) {
		input := struct {
			Filename string `json:"filename"`
		}{
			Filename: "project-1760900859010-20230128_100359.jpg",
		}

		// ejecutamos la herramienta mcp
		params := &mcp.CallToolParams{
			Name:      "get_image_project",
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
	})
}
