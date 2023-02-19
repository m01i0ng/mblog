package errno

import "fmt"

type Errno struct {
	HTTP    int
	Code    string
	Message string
}

func (e *Errno) Error() string {
	return e.Message
}

func (e *Errno) SetMessage(format string, args ...interface{}) *Errno {
	e.Message = fmt.Sprintf(format, args)
	return e
}

func Decode(err error) (int, string, string) {
	if err == nil {
		return OK.HTTP, OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Errno:
		return typed.HTTP, typed.Code, typed.Message
	default:

	}

	return InternalServerError.HTTP, InternalServerError.Code, err.Error()
}
