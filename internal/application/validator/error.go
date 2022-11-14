package validator

func NewValidationError() *ValidationError {
	return &ValidationError{
		Message: "Validation error",
		Errors:  make(map[string][]string),
	}
}

type ValidationError struct {
	Message string
	Errors  map[string][]string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func (e *ValidationError) AddError(field string, message string) {
	e.Errors[field] = append(e.Errors[field], message)
}

func (e *ValidationError) AddErrorIf(condition bool, field string, message string) {
	if condition {
		e.AddError(field, message)
	}
}

func (e *ValidationError) HasErrors() bool {
	return len(e.Errors) > 0
}
