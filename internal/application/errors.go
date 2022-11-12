package application

type ApplicationError struct {
	Err     error
	Message string
}

func NewApplicationError(err error, message string) *ApplicationError {
	return &ApplicationError{Err: err, Message: message}
}

func (e *ApplicationError) Error() string {
	return e.Message
}

type BookNotFoundError struct {
	Err     error
	Message string
}

func NewBookNotFoundError(err error, message string) *BookNotFoundError {
	return &BookNotFoundError{Err: err, Message: message}
}

func (e *BookNotFoundError) Error() string {
	return e.Message
}
