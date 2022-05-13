package request

type UpdateBook struct {
	Id          int    `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required,number"`
	Rating      int    `json:"rating" binding:"required,number"`
}
