package model

type GetCakesUsecaseParam struct {
	Page     int
	PageSize int
}

type GetCakesResponse struct {
	Meta MetaPagination `json:"meta"`
	Data []CakeResponse `json:"cakes"`
}

type MetaPagination struct {
	PageCount int   `json:"page_count"`
	TotalData int64 `json:"total_data"`
}

type CakeResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Rating      float32 `json:"rating"`
	Image       *string `json:"image"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type CakeMutationResponse struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Rating      float32 `json:"rating"`
	Image       *string `json:"image"`
}

type CakeDeleteResponse struct {
	ID int `json:"id"`
}
