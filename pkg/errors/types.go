package errors

import "net/http"

func ErrUnauthorized() *Error {
	return New(http.StatusUnauthorized, "Unauthorized")
}

func ErrInternalServer() *Error {
	return New(http.StatusInternalServerError, "服务内部错误")
}

func ErrBadRequest() *Error {
	return New(http.StatusBadRequest, "bad request")
}

func ErrNotFound() *Error {
	return New(http.StatusBadRequest, "not found")
}

func ErrInvalidRequestParams() *Error {
	return New(http.StatusUnprocessableEntity, "请求参数解析失败, 请查看errs中的提示信息")
}
