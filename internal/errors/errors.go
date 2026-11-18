package errors

type AppError struct {
	StatusCode int    `json:"_"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (e *AppError) Error() string {

	return e.Message
}

func NewAppError(status int, code string, message string) *AppError {
	return &AppError{
		StatusCode: status,
		Code:       code,
		Message:    message,
	}
}

// define the error standard
var (
	ErrInternalServer = NewAppError(500, "INTERNAL_SERVER_ERROR", "internal server error")
	ErrBadRequest     = NewAppError(400, "BAD_REQUEST_ERROR", "bad request, check the request")
	ErrUnauthorized   = NewAppError(401, "UNAUTHORIZED_ERROR", "unauthorized, please login first")
	ErrNotFound       = NewAppError(404, "NOT_FOUND_ERROR", "not found")
)
