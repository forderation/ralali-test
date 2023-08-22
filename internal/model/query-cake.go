package model

type GetCakesQuery struct {
	Limit  int
	Offset int
}

type CakePayloadQuery struct {
	Title       string
	Description *string
	Rating      float32
	Image       *string
}
