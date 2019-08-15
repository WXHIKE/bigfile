//  Copyright 2019 The bigfile Authors. All rights reserved.
//  Use of this source code is governed by a MIT-style
//  license that can be found in the LICENSE file.

package models

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/bigfile/bigfile/config"
	"github.com/bigfile/bigfile/internal/util"
)

// ChunkSize represent chunk size, default: 1MB
const ChunkSize = 1 << 20

// Chunk represents every chunk of file
type Chunk struct {
	ID        uint64    `gorm:"type:BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT;primary_key"`
	Size      int       `gorm:"type:int;column:size"`
	Hash      string    `gorm:"type:CHAR(64) NOT NULL;UNIQUE;column:hash"`
	CreatedAt time.Time `gorm:"type:TIMESTAMP(6) NOT NULL;DEFAULT:CURRENT_TIMESTAMP(6);column:createdAt"`
	UpdatedAt time.Time `gorm:"type:TIMESTAMP(6) NOT NULL;DEFAULT:CURRENT_TIMESTAMP(6);column:updatedAt"`
}

// TableName represent table name
func (c Chunk) TableName() string {
	return "chunks"
}

// Path represent the actual storage path
func (c Chunk) Path(rootPath *string) string {

	if rootPath == nil {
		rootPath = &config.DefaultConfig.Chunk.RootPath
	}
	if c.ID < 10000 {
		panic(errors.New("invalid chunk id"))
	}
	idStr := strconv.FormatUint(c.ID, 10)
	parts := make([]string, (len(idStr)/3)+1)
	index := 0
	for ; len(idStr) > 3; index++ {
		parts[index] = util.SubStrFromToEnd(idStr, -3)
		idStr = util.SubStrFromTo(idStr, 0, -3)
	}
	parts[index] = idStr
	parts = parts[1:]
	util.ReverseSlice(parts)

	dir := fmt.Sprintf("%s/%s", strings.TrimSuffix(*rootPath, "/"), filepath.Join(parts...))

	if !util.IsDir(dir) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			panic(err)
		}
	}

	return fmt.Sprintf("%s/%s", dir, strconv.FormatUint(c.ID, 10))
}
