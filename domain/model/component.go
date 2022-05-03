package model

type Component struct {
	ID       uint   `json:"componentId"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	Retry    int    `json:"retry"`
	IsHealth bool   `json:"isHealth"`
}
