package errs

type NotFoundUser struct {
	Message string
}

func (e *NotFoundUser) Error() string {
	return "Not found"
}
