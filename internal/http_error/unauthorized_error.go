package http_error

type UnauthorizedError struct{}

func (e UnauthorizedError) Error() string {
	return "Unauthorized."
}
