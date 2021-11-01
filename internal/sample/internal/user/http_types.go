package user

// BaseResponse defines basic model for response
type BaseResponse struct {
	Message string `json:"message"`
}

type BaseListResponse struct {
	BaseResponse
	Page    int `json:"page"`
	Perpage int `json:"perpage"`
	Total   int `json:"total"`
}

type UserResponse struct {
	BaseResponse
	Data User `json:"data"`
}

type UserListResponse struct {
	BaseListResponse
	Data []User `json:"data"`
}
