// Copyright 2023 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package store

import (
	"context"
	"errors"

	"github.com/m01i0ng/mblog/internal/pkg/model"
	"gorm.io/gorm"
)

type UserStore interface {
	Create(ctx context.Context, user *model.UserM) error
	Get(ctx context.Context, username string) (*model.UserM, error)
	Update(ctx context.Context, user *model.UserM) error
	List(ctx context.Context, offset, limit int) (int64, []*model.UserM, error)
	Delete(ctx context.Context, username string) error
}

type users struct {
	db *gorm.DB
}

func (u *users) List(ctx context.Context, offset, limit int) (count int64, users []*model.UserM, err error) {
	err = u.db.Offset(offset).Limit(defaultLimit(limit)).Order("id desc").Find(&users).
		Offset(-1).
		Limit(-1).
		Count(&count).
		Error
	return
}

func (u *users) Delete(ctx context.Context, username string) error {
	err := u.db.Where("username = ?", username).Delete(&model.UserM{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}

func newUsers(db *gorm.DB) *users {
	return &users{db}
}

var _ UserStore = (*users)(nil)

func (u *users) Create(ctx context.Context, user *model.UserM) error {
	return u.db.Create(&user).Error
}

func (u *users) Get(ctx context.Context, username string) (*model.UserM, error) {
	var user model.UserM
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *users) Update(ctx context.Context, user *model.UserM) error {
	return u.db.Save(user).Error
}
