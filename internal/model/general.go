package model

type ErrorResponse struct {
	Err            error
	HttpStatusCode int
	ErrData        interface{}
}

type ErrorDetailResponse struct {
	Detail string `json:"detail"`
}
