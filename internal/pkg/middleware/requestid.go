// Copyright 2023 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/m01i0ng/mblog/internal/pkg/known"
)

// RequestID 是一个 Gin 中间件，在每个 http 请求的 context response 中注入 ID.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(known.XRequestIDKey)
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set(known.XRequestIDKey, requestID)
		c.Writer.Header().Set(known.XRequestIDKey, requestID)
		c.Next()
	}
}
