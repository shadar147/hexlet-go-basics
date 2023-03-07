package ex30

import "errors"

// некритичная ошибка валидации
type NonCriticalError struct{}

func (e NonCriticalError) Error() string {
	return "validation error"
}

// критичные ошибки
var (
	ErrBadConnection = errors.New("bad connection")
	ErrBadRequest    = errors.New("bad request")
)

const UnknownErrorMsg = "unknown error"

func GetErrorMsg(err error) string {
	if errors.Is(err, ErrBadConnection) || errors.Is(err, ErrBadRequest) {
		return err.Error()
	}

	if errors.As(err, &NonCriticalError{}) {
		return ""
	}

	return UnknownErrorMsg
}
