// Copyright 2022 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package mblog

import (
	"encoding/json"
	"fmt"

	"github.com/m01i0ng/mblog/internal/pkg/log"
	"github.com/m01i0ng/mblog/pkg/version/verflag"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func NewMBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "mblog",
		Short:        "A good Go project",
		Long:         "",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Init(logOptions())
			defer log.Sync()
			verflag.PrintAndExitIfRequested()
			return run()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), arg)
				}
			}
			return nil
		},
	}
	cobra.OnInitialize(initConfig)
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the config file. Empty string for no config file.")
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	verflag.AddFlags(cmd.PersistentFlags())
	return cmd
}

func run() error {
	settings, _ := json.Marshal(viper.AllSettings())
	log.Infow(string(settings))
	log.Infow(viper.GetString("db.username"))
	return nil
}
