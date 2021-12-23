package domain

type NotFoundError struct {
	mgs string
}

func (e *NotFoundError) Error() string {
	return e.mgs
}

func NewNotFoundError(msg string) *NotFoundError {
	return &NotFoundError{msg}
}
