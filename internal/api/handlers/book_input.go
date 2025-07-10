package handlers

type CreateBookInput struct {
	Title  string  `json:"title" binding:"required"`
	Author string  `json:"author" binding:"required"`
	Price  float64 `json:"price" binding:"required,gt=0"`
	ISBN   string  `json:"isbn" binding:"required"`
}

type UpdateBookInput struct {
	Title  string  `json:"title" binding:"required"`
	Author string  `json:"author" binding:"required"`
	Price  float64 `json:"price" binding:"required,gt=0"`
	ISBN   string  `json:"isbn" binding:"required"`
}
