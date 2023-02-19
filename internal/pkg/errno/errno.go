// Copyright 2023 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

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
