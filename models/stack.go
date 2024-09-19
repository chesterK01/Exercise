package models

type Stack struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Stock   int    `json:"stock"`
	Quality string `json:"quality"`
}
type UpdateStockRequest struct {
	Stock int `json:"stock"` // Struct chứa thông tin cập nhật stock
}
type UpdateQualityRequest struct {
	Quality string `json:"quality"`
}
