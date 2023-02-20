// Copyright 2023 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package errno

var (
	OK = &Errno{
		HTTP:    200,
		Code:    "",
		Message: "",
	}

	InternalServerError = &Errno{
		HTTP:    500,
		Code:    "InternalError",
		Message: "Internal server error.",
	}

	ErrPageNotFound = &Errno{
		HTTP:    404,
		Code:    "ResourceNotFound.PageNotFound",
		Message: "Page not found.",
	}

	ErrBind = &Errno{HTTP: 400, Code: "InvalidParameter.BindError", Message: "Error occurred while binding the request body to the struct."}

	ErrInvalidParameter = &Errno{HTTP: 400, Code: "InvalidParameter", Message: "Parameter verification failed."}
)
