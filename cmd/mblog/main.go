// Copyright 2022 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package main

import (
	"os"

	"github.com/m01i0ng/mblog/internal/mblog"
)

func main() {
	command := mblog.NewMBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
