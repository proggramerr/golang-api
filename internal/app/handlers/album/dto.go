package album

type CreateAlbumInput struct {
	Title  string   `json:"title" binding:"required"`
	Artist string   `json:"artist" binding:"required"`
	Price  *float64 `json:"price" binding:"required"`
}

type DeleteAlbumInput struct {
	ID uint `json:"id"`
}
