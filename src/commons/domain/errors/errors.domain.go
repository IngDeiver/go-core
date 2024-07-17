package errorsDomain

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrUnauthorizedError = errors.New("unauthorized")
	ErrNotFoundError = errors.New("not found resource")
	ErrUserAlreadyExistsError = errors.New("user already exists")
	ErrInvalidTokenError = errors.New("invalid token")
)