package application

func NewApplicationError(err error, message string) *ApplicationError {
	return &ApplicationError{Err: err, Message: message}
}

type ApplicationError struct {
	Err     error
	Message string
}

func (e *ApplicationError) Error() string {
	return e.Message
}
