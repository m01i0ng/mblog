// Copyright 2022 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package mblog

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m01i0ng/mblog/internal/mblog/controller/v1/user"
	"github.com/m01i0ng/mblog/internal/mblog/store"
	"github.com/m01i0ng/mblog/internal/pkg/known"
	"github.com/m01i0ng/mblog/internal/pkg/log"
	"github.com/m01i0ng/mblog/internal/pkg/middleware"
	pb "github.com/m01i0ng/mblog/pkg/proto/mblog/v1"
	"github.com/m01i0ng/mblog/pkg/token"
	"github.com/m01i0ng/mblog/pkg/version/verflag"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc"
)

var cfgFile string

func NewMBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "mblog.sql",
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

func startGRPCServer() *grpc.Server {
	lis, err := net.Listen("tcp", viper.GetString("grpc.addr"))
	if err != nil {
		log.Fatalw("Failed to listen", "err", err)
	}

	server := grpc.NewServer()
	pb.RegisterMBlogServer(server, user.New(store.S, nil))

	log.Infow("Start to listening the incoming requests on grpc address", "addr", viper.GetString("grpc.addr"))
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalw(err.Error())
		}
	}()

	return server
}

func startInsecureServer(g *gin.Engine) *http.Server {
	httpSrv := &http.Server{
		Addr:    viper.GetString("addr"),
		Handler: g,
	}

	log.Infow("Start to listen the incoming requests on http address", "addr", viper.GetString("addr"))
	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()

	return httpSrv
}

func startSecureServer(g *gin.Engine) *http.Server {
	httpsSrv := &http.Server{
		Addr:    viper.GetString("tls.addr"),
		Handler: g,
	}

	log.Infow("Start to listen the incoming requests on https address", "addr", viper.GetString("tls.addr"))
	cert, key := viper.GetString("tls.cert"), viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			if err := httpsSrv.ListenAndServeTLS(cert, key); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalw(err.Error())
			}
		}()
	}

	return httpsSrv
}

func run() error {
	if err := initStore(); err != nil {
		return err
	}

	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	mws := []gin.HandlerFunc{gin.Recovery(), middleware.NoCache, middleware.Cors, middleware.Secure, middleware.RequestID()}
	g.Use(mws...)

	if err := installRouters(g); err != nil {
		return err
	}

	token.Init(viper.GetString("jwt-secret"), known.XUsernameKey)

	httpSrv := startInsecureServer(g)
	grpcServer := startGRPCServer()

	quit := make(chan os.Signal)
	signal.Notify(quit, unix.SIGINT, unix.SIGTERM)
	<-quit
	log.Infow("Shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Errorw("Insecure Server forced to shutdown", "err", err)
		return err
	}
	grpcServer.GracefulStop()

	log.Infow("Server exiting")

	return nil
}
