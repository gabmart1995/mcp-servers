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

type Project struct {
	Name        string `json:"name"`
	Id          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Created_at  string `json:"created_at"`
	Image       string `json:"image"`
}
