// Copyright 2023 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package store

import (
	"sync"

	"gorm.io/gorm"
)

var (
	once sync.Once
	S    *datastore
)

type IStore interface {
	DB() *gorm.DB
	Users() UserStore
}

type datastore struct {
	db *gorm.DB
}

func (d *datastore) DB() *gorm.DB {
	return d.db
}

func NewStore(db *gorm.DB) *datastore {
	once.Do(func() {
		S = &datastore{db}
	})
	return S
}

func (d *datastore) Users() UserStore {
	return newUsers(d.db)
}

var _ IStore = (*datastore)(nil)
