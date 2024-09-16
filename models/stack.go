package models

type Stack struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Stock   int    `json:"stock"`
	Quality string `json:"quality"`
}
