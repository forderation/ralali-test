package model

type RepoGetCakesParam struct {
	Limit  int
	Offset int
}

type RepoCakeParam struct {
	Title       string
	Description *string
	Rating      float32
	Image       *string
}
