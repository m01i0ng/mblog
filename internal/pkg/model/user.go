// Copyright 2023 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package model

import (
	"time"

	"github.com/m01i0ng/mblog/internal/pkg/log"
	"github.com/m01i0ng/mblog/pkg/auth"
	"gorm.io/gorm"
)

type UserM struct {
	CreatedAt time.Time `gorm:"column:createdAt"`      //
	Email     string    `gorm:"column:email"`          //
	ID        int64     `gorm:"column:id;primary_key"` //
	Nickname  string    `gorm:"column:nickname"`       //
	Password  string    `gorm:"column:password"`       //
	Phone     string    `gorm:"column:phone"`          //
	UpdatedAt time.Time `gorm:"column:updatedAt"`      //
	Username  string    `gorm:"column:username"`       //
}

// TableName sets the insert table name for this struct type.
func (n *UserM) TableName() string {
	return "user"
}

// BeforeCreate 在创建数据库记录之前加密明文密码.
func (n *UserM) BeforeCreate(tx *gorm.DB) (err error) {
	n.Password, err = auth.Encrypt(n.Password)
	log.Infow("New user", n)
	if err != nil {
		return err
	}

	return nil
}
