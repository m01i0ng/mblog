// Copyright 2023 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/m01i0ng/mblog/internal/pkg/log"
	pb "github.com/m01i0ng/mblog/pkg/proto/mblog/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr  = flag.String("addr", "localhost:9090", "")
	limit = flag.Int64("limit", 10, "")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalw("Didn't connect", "err", err)
	}
	defer conn.Close()

	c := pb.NewMBlogClient(conn)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()

	var offset int64 = 0

	r, err := c.ListUser(ctx, &pb.ListUserRequest{
		Limit:  limit,
		Offset: &offset,
	})
	if err != nil {
		log.Fatalw("Couldn't greet", "err", err)
	}

	fmt.Println("TotalCount: ", r.TotalCount)

	for _, user := range r.Users {
		s, _ := json.Marshal(user)
		fmt.Println(s)
	}
}
