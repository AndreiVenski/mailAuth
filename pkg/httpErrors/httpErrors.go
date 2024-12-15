package httpErrors

import "github.com/pkg/errors"

var (
	InvalidRequestDataError = errors.New("Invalid request data")
	ExistedUserError        = errors.New("User just exists")
	NotExistedCodeError     = errors.New("Not found code")
)

func IsUserError(err error) bool {
	switch {
	case errors.Is(err, InvalidRequestDataError), errors.Is(err, ExistedUserError), errors.Is(err, NotExistedCodeError):
		return true
	default:
		return false
	}
}
