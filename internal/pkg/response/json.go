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

func NewServerErrorResponse(err error) *JsonResponse {
	return &JsonResponse{
		Code:    1,
		Message: err.Error(),
		Data:    nil,
	}
}
