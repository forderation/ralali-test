package model

type ErrorResponse struct {
	Err            error
	HttpStatusCode int
	ErrData        interface{}
}

type ErrorDetailResponse struct {
	Detail string `json:"detail"`
}

type JsonErrorResp struct {
	ErrorMessage string      `json:"error_message"`
	ErrData      interface{} `json:"error_data,omitempty"`
}
