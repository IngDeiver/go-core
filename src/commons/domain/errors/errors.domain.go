package errorsDomain

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrUnauthorizedError = errors.New("unauthorized")
	ErrNotFoundError = errors.New("not found resource")
)