// Copyright 2023 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package version

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/gosuri/uitable"
)

var (
	GitVersion   = "v0.0.0-master+$Format:%h$"
	BuildDate    = "1970-01-01T00:00:00Z"
	GitCommit    = "$Format:%H$"
	GitTreeState = ""
)

type Info struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func (i Info) Text() ([]byte, error) {
	table := uitable.New()
	table.RightAlign(0)
	table.MaxColWidth = 80
	table.Separator = " "
	table.AddRow("gitVersion:", i.GitVersion)
	table.AddRow("gitCommit:", i.GitCommit)
	table.AddRow("gitTreeState:", i.GitTreeState)
	table.AddRow("buildDate:", i.BuildDate)
	table.AddRow("goVersion:", i.GoVersion)
	table.AddRow("compiler:", i.Compiler)
	table.AddRow("platform:", i.Platform)

	return table.Bytes(), nil
}

func (i Info) String() string {
	if s, err := i.Text(); err == nil {
		return string(s)
	}
	return i.GitVersion
}

func (i Info) ToJSON() string {
	s, _ := json.Marshal(i)
	return string(s)
}

func Get() Info {
	return Info{
		GitVersion:   GitVersion,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildDate:    BuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
