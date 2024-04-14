package custom_errors

import (
	"errors"
)

var ErrUserHasNotAccess = errors.New("User does not have access")
var ErrUserUnAuthorized = errors.New("User is not authorized")
var ErrStatusNotFound = errors.New("Banner for not found")
