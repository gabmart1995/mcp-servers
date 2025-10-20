package main

import (
	"context"
	"log"
	"mcp-projects-go/controllers"
	"mcp-projects-go/models"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

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
		controllers.SaveProject,
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
		controllers.ListProjects,
	)

	server.AddTool(
		&mcp.Tool{
			Name:        "list_project_id",
			Description: "Lista un proyecto en especifico usando un ID",
			InputSchema: models.InputSchema{
				Type: "object",
				Properties: map[string]models.Properties{
					"id": {Type: "string", Description: "Identificador del proyecto"},
				},
				Required: []string{"id"},
			},
		},
		controllers.ListProjectId,
	)

	server.AddTool(
		&mcp.Tool{
			Name:        "update_project",
			Description: "Actualiza un proyecto pasando el identificador",
			InputSchema: models.InputSchema{
				Type: "object",
				Properties: map[string]models.Properties{
					"id":          {Type: "string", Description: "identificador del proyecto"},
					"name":        {Type: "string", Description: "nombre del proyecto"},
					"description": {Type: "string", Description: "descripcion"},
					"state":       {Type: "string", Description: "estado del proyecto"},
				},
				Required: []string{"id", "name", "description", "state"},
			},
		},
		controllers.UpdateProject,
	)

	server.AddTool(
		&mcp.Tool{
			Name:        "delete_project_id",
			Description: "Elimina un proyecto en especifico usando un ID",
			InputSchema: models.InputSchema{
				Type: "object",
				Properties: map[string]models.Properties{
					"id": {Type: "string", Description: "Identificador del proyecto"},
				},
				Required: []string{"id"},
			},
		},
		controllers.DeleteProject,
	)

	server.AddTool(
		&mcp.Tool{
			Name:        "get_image_project",
			Description: "Localiza una imagen del proyecto usando el nombre del archivo",
			InputSchema: models.InputSchema{
				Type: "object",
				Properties: map[string]models.Properties{
					"filename": {Type: "string", Description: "nombre del archivo"},
				},
				Required: []string{"filename"},
			},
		},
		controllers.GetImageProject,
	)

	// corremos el server
	log.Fatal(server.Run(context.Background(), &mcp.StdioTransport{}))
}
