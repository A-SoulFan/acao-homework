package response

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessJsonResponse(data interface{}) *JsonResponse {
	return &JsonResponse{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}

// NewServerErrorResponse TODO: 后续实现类型区分
func NewServerErrorResponse(err error) *JsonResponse {
	return &JsonResponse{
		Code:    1,
		Message: err.Error(),
		Data:    nil,
	}
}
