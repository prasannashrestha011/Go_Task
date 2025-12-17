package utils

//custom error type
// status->true(success)
// status->false(failed)
type AppError struct {
	Message string
	Err error
	Code int
}

func (e *AppError) Error() string{
	return e.Message
}

func New(code int,message string,err error) *AppError{

	return &AppError{
		Message: message,
		Code: code,
		Err:err,
		}
}

var (
	ErrBadRequest=New(400,"Invalid request body",nil)
	ErrUnAuthorized=New(401,"Request unauthorized",nil)
)
