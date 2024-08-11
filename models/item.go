package models

type Item struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Url string `json:"url"`
	Description string `json:"description"`
}
