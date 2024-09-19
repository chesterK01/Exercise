package models

type Stack struct {
	ID      int    `json:"id"`
	BookID  int    `json:"book_id"` // ID của sách, là khóa ngoại từ bảng author_book
	Stock   int    `json:"stock"`
	Quality string `json:"quality"`
}

type UpdateStockQualityRequest struct {
	Stock   int    `json:"stock"`
	Quality string `json:"quality"`
}
