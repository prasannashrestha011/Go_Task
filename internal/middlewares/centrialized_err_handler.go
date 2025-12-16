package middlewares

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func BadRequest(msg string) *AppError {
	return &AppError{
		Code:    400,
		Message: msg,
	}
}
func InternalServerError(msg string) *AppError {
	return &AppError{
		Code:    500,
		Message: msg,
	}
}

func NotFoundError(msg string) *AppError {
	return &AppError{
		Code:    404,
		Message: msg,
	}
}
func ForbiddenError(msg string) *AppError {
	return &AppError{
		Code:    403,
		Message: msg,
	}
}

func BadGatewayError(msg string) *AppError {
	return &AppError{
		Code:    502,
		Message: msg,
	}
}
