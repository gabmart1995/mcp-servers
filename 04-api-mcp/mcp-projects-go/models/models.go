package models

type InputSchema struct {
	Type       string                `json:"type"`
	Required   []string              `json:"required"`
	Properties map[string]Properties `json:"properties"`
}

type Properties struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
