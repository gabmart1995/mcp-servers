package controllers

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const API_URL = "http://localhost:3000/project"

// salva un nuevo projecto
func SaveProject(ctx context.Context, ctr *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
func ListProjects(ctx context.Context, ctr *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	request, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/list", API_URL),
		nil,
	)

	if err != nil {
		return nil, err
	}

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

// lista un proyecto usando su identificador
func ListProjectId(ctx context.Context, ctr *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var args struct {
		Id string `json:"id"`
	}

	if err := json.Unmarshal(ctr.Params.Arguments, &args); err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/list/%s", API_URL, args.Id), nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close() // close the buffer

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

func UpdateProject(ctx context.Context, ctr *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	data, err := json.Marshal(ctr.Params.Arguments)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/update", API_URL),
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

	defer response.Body.Close() // close the buffer

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

func DeleteProject(ctx context.Context, ctr *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var args struct {
		Id string `json:"id"`
	}

	if err := json.Unmarshal(ctr.Params.Arguments, &args); err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/delete/%s", API_URL, args.Id),
		nil,
	)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close() // close the buffer

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

func GetImageProject(ctx context.Context, ctr *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var args struct {
		Filename string `json:"filename"`
	}
	formats := map[string]string{
		"jpg":  "image/jpeg",
		"jpeg": "image/jpeg",
		"git":  "image/gif",
		"png":  "image/png",
	}

	if err := json.Unmarshal(ctr.Params.Arguments, &args); err != nil {
		return nil, err
	}

	// validamos la extension antes de realizar la consulta
	extension := filepath.Ext(args.Filename)[1:]
	mimeType, exists := formats[extension]

	if !exists {
		return nil, errors.New("error: format not supported")
	}

	request, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/image/%s", API_URL, args.Filename),
		nil,
	)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close() // close the buffer

	// extraemos el body
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	// construimos la instacia de la imagen
	imageBase64 := base64.StdEncoding.EncodeToString(body)

	result := &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.ImageContent{
				MIMEType: mimeType,
				Data:     []byte(imageBase64),
			},
		},
	}

	return result, nil
}
