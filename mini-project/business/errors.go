package business

type ErrDuplicateEntry struct{}

func (e *ErrDuplicateEntry) Error() string {
	return "data already exists"
}

type ErrNotFound struct{}

func (e *ErrNotFound) Error() string {
	return "data not found"
}

type ErrUnauthorized struct {
	Message string
}

func (e *ErrUnauthorized) Error() string {
	if e.Message == "" {
		e.Message = "unauthorized"
	}
	return e.Message
}

func NewErrUnauthorized(msg string) *ErrUnauthorized {
	return &ErrUnauthorized{msg}
}
