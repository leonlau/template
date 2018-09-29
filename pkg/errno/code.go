package errno

var (
	// Common error
	OK                  = &Errno{Code: 0, Message: "OK"}
	ErrNotFound         = &Errno{Code: 4040, Message: "Not found!"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// user errors
	ErrValidation   = &Errno{Code: 20001, Message: "Validation failed."}
	ErrTokenInvalid = &Errno{Code: 20103, Message: "The token was invalid."}

	ErrPasswordIncorrect = &Errno{Code: 4000, Message: "密码错误"}
	ErrToken             = &Errno{Code: 5000, Message: "token 错误"}
)
