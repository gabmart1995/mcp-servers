package main

import (
	"context"
	"log"
	"os/exec"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestGetImage(t *testing.T) {
	ctx := context.Background()
	input := struct {
		Filename string `json:"filename"`
	}{Filename: "project-1760900859010-20230128_100359.jpg"}

	// creamos el cliente mcp
	client := mcp.NewClient(&mcp.Implementation{Name: "test-mcp-project", Version: "0.0.1"}, nil)

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
		Name:      "get_image_project",
		Arguments: input,
	}

	// evaluamos la expresion obtenida
	t.Run("get_image", func(t *testing.T) {
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
