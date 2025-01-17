//  Copyright 2019 The bigfile Authors. All rights reserved.
//  Use of this source code is governed by a MIT-style
//  license that can be found in the LICENSE file.

package models

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/bigfile/bigfile/internal/util"
	"github.com/jinzhu/gorm"
)

// IsDir represent a file model is a directory actually
const IsDir = int8(1)

// Hidden represent some file is hidden
const Hidden = int8(1)

var (
	// ErrFileExisted represent that the path has been occupied
	ErrFileExisted = errors.New("file has already existed")
	// ErrOverwriteDir represent that try to overwrite some directory
	ErrOverwriteDir = errors.New("directory can't be overwritten")
	// ErrAppendToDir represent that try to append content to directory
	ErrAppendToDir = errors.New("can't append data to directory")
	// ErrReadDir represent that can't read data from directory, only file
	ErrReadDir = errors.New("can't read a directory")
	// ErrAccessDenied represent a file can't be accessed by some tokens
	ErrAccessDenied = errors.New("file can't be accessed by some tokens")
	// ErrDeleteNonEmptyDir represent delete non-empty directory
	ErrDeleteNonEmptyDir = errors.New("delete non-empty directory")
)

// File represent a file or a directory of system. If it's a file
// it has to associate with an object. Actually, the object hold
// the real content of file.
type File struct {
	ID            uint64     `gorm:"type:BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT;primary_key"`
	UID           string     `gorm:"type:CHAR(32) NOT NULL;UNIQUE;column:uid"`
	PID           uint64     `gorm:"type:BIGINT(20) UNSIGNED NOT NULL;column:pid"`
	AppID         uint64     `gorm:"type:BIGINT(20) UNSIGNED NOT NULL;column:appId"`
	ObjectID      uint64     `gorm:"type:BIGINT(20) UNSIGNED NOT NULL;column:objectId"`
	Size          int        `gorm:"type:int;column:size"`
	Name          string     `gorm:"type:VARCHAR(255);NOT NULL;column:name"`
	Ext           string     `gorm:"type:VARCHAR(255);NOT NULL;column:ext"`
	IsDir         int8       `gorm:"type:tinyint;column:isDir;DEFAULT:0"`
	Hidden        int8       `gorm:"type:tinyint;column:hidden;DEFAULT:0"`
	DownloadCount uint64     `gorm:"type:BIGINT(20);column:downloadCount;DEFAULT:0"`
	CreatedAt     time.Time  `gorm:"type:TIMESTAMP(6) NOT NULL;DEFAULT:CURRENT_TIMESTAMP(6);column:createdAt"`
	UpdatedAt     time.Time  `gorm:"type:TIMESTAMP(6) NOT NULL;DEFAULT:CURRENT_TIMESTAMP(6);column:updatedAt"`
	DeletedAt     *time.Time `gorm:"type:TIMESTAMP(6);INDEX;column:deletedAt"`

	App       App       `gorm:"foreignkey:appId;association_autoupdate:false;association_autocreate:false"`
	Object    Object    `gorm:"foreignkey:objectId;association_autoupdate:false;association_autocreate:false"`
	Parent    *File     `gorm:"foreignkey:id;association_foreignkey:pid;association_autoupdate:false;association_autocreate:false"`
	Children  []File    `gorm:"foreignkey:pid;association_foreignkey:id;association_autoupdate:false;association_autocreate:false"`
	Histories []History `gorm:"foreignkey:fileId;association_autoupdate:false;association_autocreate:false"`
}

func pathCacheKey(app *App, path string) string {
	return fmt.Sprintf("%d-%s", app.ID, path)
}

// TableName represent the name of files table
func (f *File) TableName() string {
	return "files"
}

func (f *File) executeDelete(forceDelete bool, db *gorm.DB) error {
	if f.IsDir == 0 {
		return db.Delete(f).Error
	}

	var err error

	if err = db.Preload("Children").Find(f).Error; err != nil {
		return err
	}

	if len(f.Children) == 0 {
		return db.Delete(f).Error
	}

	if forceDelete {
		for _, child := range f.Children {
			if err = child.executeDelete(forceDelete, db); err != nil {
				return err
			}
		}
		db.Model(f).Update("size", 0)
		return db.Delete(f).Error
	}
	return ErrDeleteNonEmptyDir
}

// Delete is used to delete file or directory. if the file is a non-empty directory,
// 'forceDelete' determine to delete or not sub directories and files
func (f *File) Delete(forceDelete bool, db *gorm.DB) (err error) {

	if f.Parent == nil {
		if err = db.Preload("Parent").Find(f).Error; err != nil {
			return err
		}
	}

	originSize := f.Size
	if err = f.executeDelete(forceDelete, db); err != nil {
		return err
	}

	if originSize != 0 {
		if err = f.Parent.UpdateParentSize(-originSize, db); err != nil {
			return err
		}
	}

	return db.Unscoped().Find(f).Error
}

// CanBeAccessedByToken represent whether the file can be accessed by the token
func (f *File) CanBeAccessedByToken(token *Token, db *gorm.DB) error {
	var (
		err  error
		path string
	)
	if path, err = f.Path(db); err != nil {
		return err
	}
	if !strings.HasPrefix(path, token.Path) {
		return ErrAccessDenied
	}
	return nil
}

// Reader is used to get reader that continues to read data from underlying
// chunk until io.EOF
func (f *File) Reader(rootPath *string, db *gorm.DB) (io.ReadSeeker, error) {
	if f.IsDir == 1 {
		return nil, ErrReadDir
	}
	if f.Object.ID == 0 {
		db.Preload("Object").Find(&f)
	}
	return (&f.Object).Reader(rootPath, db)
}

// Path is used to get the complete path of file
func (f *File) Path(db *gorm.DB) (string, error) {

	if f.PID == 0 && f.IsDir == IsDir {
		return "/", nil
	}

	var (
		parts   []string
		current = *f
	)
	for {
		parts = append(parts, current.Name)
		if current.PID == 0 {
			break
		}
		temp := &File{}
		db.Where("id = ?", current.PID).Find(temp)
		current = *temp
	}

	util.ReverseSlice(parts)

	return strings.Join(parts, "/"), nil
}

// UpdateParentSize is used to update parent size. note, size may be a negative number.
func (f *File) UpdateParentSize(size int, db *gorm.DB) error {
	f.Size += size
	if err := db.Model(f).Update("size", f.Size).Error; err != nil {
		return err
	}
	if f.PID == 0 {
		return nil
	}
	if f.Parent == nil {
		f.Parent = &File{}
	}
	if err := db.Model(f).Association("Parent").Find(f.Parent).Error; err != nil {
		return err
	}
	return f.Parent.UpdateParentSize(size, db)
}

func (f *File) createHistory(objectID uint64, path string, db *gorm.DB) error {
	return db.Save(&History{ObjectID: objectID, FileID: f.ID, Path: path}).Error
}

// OverWriteFromReader is used to overwrite the object
func (f *File) OverWriteFromReader(reader io.Reader, hidden int8, rootPath *string, db *gorm.DB) (err error) {

	if f.IsDir == IsDir {
		return ErrOverwriteDir
	}

	var (
		path     string
		object   *Object
		sizeDiff int
	)

	if path, err = f.Path(db); err != nil {
		return err
	}

	if err := f.createHistory(f.ObjectID, path, db); err != nil {
		return err
	}

	if object, err = CreateObjectFromReader(reader, rootPath, db); err != nil {
		return err
	}

	f.Object = *object
	f.ObjectID = object.ID
	f.Hidden = hidden
	sizeDiff = object.Size - f.Size
	f.Size += sizeDiff

	if err = db.Model(f).Update(map[string]interface{}{
		"objectId": object.ID,
		"hidden":   hidden,
		"size":     f.Size,
	}).Error; err != nil {
		return err
	}
	db.Preload("Parent").Preload("App").Find(f)
	return f.Parent.UpdateParentSize(sizeDiff, db)
}

func (f *File) mustPath(db *gorm.DB) string {
	path, _ := f.Path(db)
	return path
}

// MoveTo move file to another path, the input path must be complete and new path.
// if the input path ios the same as the previous path, nothing changes.
func (f *File) MoveTo(newPath string, db *gorm.DB) (err error) {
	var (
		newPathDir      = filepath.Dir(newPath)
		newPathDirFile  *File
		newPathFileName = filepath.Base(newPath)
		newPathExt      = strings.TrimPrefix(filepath.Ext(newPathFileName), ".")
		previousPath    string
	)

	if previousPath, err = f.Path(db); err != nil {
		return err
	}

	if previousPath == newPath {
		return nil
	}

	if f.App.ID == 0 {
		if err = db.Preload("App").Find(f).Error; err != nil {
			return err
		}
	}

	if _, err := FindFileByPathWithTrashed(&f.App, newPath, db); err == nil {
		return ErrFileExisted
	}

	if newPathDirFile, err = CreateOrGetLastDirectory(&f.App, newPathDir, db); err != nil {
		return err
	}

	if f.IsDir == 0 {
		if err = f.createHistory(f.ObjectID, previousPath, db); err != nil {
			return err
		}
	}

	f.Name = newPathFileName
	f.Ext = newPathExt

	defer func() {
		pathToFileCache.Delete(previousPath)
		_ = pathToFileCache.Add(pathCacheKey(&f.App, f.mustPath(db)), f, time.Minute*10)
	}()

	// only change the file name, still is in the same directory
	if newPathDirFile.ID == f.PID {
		return db.Model(f).Updates(map[string]interface{}{"name": f.Name, "ext": f.Ext}).Error
	}

	if f.Parent == nil || f.Parent.ID == 0 {
		f.Parent = &File{}
		if err = db.Model(f).Association("Parent").Find(f.Parent).Error; err != nil {
			return err
		}
	}

	if err = newPathDirFile.UpdateParentSize(f.Size, db); err != nil {
		return err
	}

	if err = f.Parent.UpdateParentSize(-f.Size, db); err != nil {
		return err
	}
	f.PID = newPathDirFile.ID
	f.Parent = newPathDirFile
	f.Name = newPathFileName
	f.Ext = newPathExt

	return db.Model(f).Updates(map[string]interface{}{"pid": f.PID, "name": f.Name, "ext": f.Ext}).Error
}

// AppendFromReader is used to append content from reader to file
func (f *File) AppendFromReader(reader io.Reader, hidden int8, rootPath *string, db *gorm.DB) (err error) {

	if f.IsDir == IsDir {
		return ErrAppendToDir
	}

	var (
		size   int
		object *Object
	)

	if err = db.Preload("Object").Preload("Parent").Preload("App").First(f).Error; err != nil {
		return err
	}

	if object, size, err = f.Object.AppendFromReader(reader, rootPath, db); err != nil {
		return err
	}

	f.Hidden = hidden
	f.Size += size
	f.Object = *object
	f.ObjectID = object.ID

	if err = db.Model(f).Updates(map[string]interface{}{"hidden": f.Hidden, "size": f.Size, "objectId": f.ObjectID}).Error; err != nil {
		return err
	}

	return f.Parent.UpdateParentSize(size, db)
}

// CreateOrGetLastDirectory is used to get last level directory, there is no difference
// between a relative path and an absolute path
func CreateOrGetLastDirectory(app *App, dirPath string, db *gorm.DB) (*File, error) {
	var (
		parent = &File{ID: 0}
		err    error
		parts  = strings.Split(strings.TrimRight(strings.TrimSpace(dirPath), "/"), "/")
	)

	if parts[0] != "" {
		if parent, err = CreateOrGetRootPath(app, db); err != nil {
			return nil, err
		}
	}

	for _, part := range parts {
		file := &File{}
		if err = db.Where("appId = ? and pid = ? and name = ?", app.ID, parent.ID, part).First(file).Error; err != nil {
			if !util.IsRecordNotFound(err) {
				return nil, err
			}
			file.AppID = app.ID
			file.PID = parent.ID
			file.Name = part
			file.IsDir = 1
			file.UID = UID()
			if err = db.Save(file).Error; err != nil {
				return nil, err
			}
		}
		parent = file
	}
	parent.App = *app
	return parent, nil
}

// CreateOrGetRootPath is used to create or get root directory
func CreateOrGetRootPath(app *App, db *gorm.DB) (*File, error) {
	var (
		file = &File{}
		err  error
	)
	err = db.Where("appId = ? and pid = 0 and name = ''", app.ID).First(file).Error
	file.App = *app
	return file, err
}

// CreateFileFromReader is used to create a file from reader.
func CreateFileFromReader(app *App, path string, reader io.Reader, hidden int8, rootPath *string, db *gorm.DB) (file *File, err error) {
	var (
		object    *Object
		parentDir *File
		dirPrefix = filepath.Dir(path)
		fileName  = filepath.Base(path)
	)

	if f, err := FindFileByPathWithTrashed(app, path, db); err == nil && f.ID > 0 {
		return nil, ErrFileExisted
	}

	if parentDir, err = CreateOrGetLastDirectory(app, dirPrefix, db); err != nil {
		return nil, err
	}

	if object, err = CreateObjectFromReader(reader, rootPath, db); err != nil {
		return nil, err
	}

	file = &File{
		UID:      UID(),
		PID:      parentDir.ID,
		AppID:    app.ID,
		ObjectID: object.ID,
		Size:     object.Size,
		Name:     fileName,
		Ext:      strings.TrimPrefix(filepath.Ext(fileName), "."),
		Hidden:   hidden,
		Object:   *object,
		App:      *app,
		Parent:   parentDir,
	}

	if err = db.Create(file).Error; err != nil {
		return nil, err
	}

	return file, parentDir.UpdateParentSize(object.Size, db)
}

// FindFileByUID is used to find a file by uid
func FindFileByUID(uid string, trashed bool, db *gorm.DB) (*File, error) {
	var (
		file = &File{}
		err  error
	)
	if trashed {
		db = db.Unscoped()
	}
	if err = db.Where("uid = ?", uid).Find(file).Error; err != nil {
		return file, err
	}
	return file, nil
}

// FindFileByPathWithTrashed is used to find a file by the specify path, include deleted path
func FindFileByPathWithTrashed(app *App, path string, db *gorm.DB) (*File, error) {
	return FindFileByPath(app, path, db.Unscoped(), true)
}

// FindFileByPath is used to find a file by the specify path
func FindFileByPath(app *App, path string, db *gorm.DB, useCache bool) (*File, error) {
	var cacheKey = pathCacheKey(app, path)

	if useCache {
		if fileValue, ok := pathToFileCache.Get(cacheKey); ok {
			file := fileValue.(*File)
			if err := db.Where("id = ?", file.ID).Find(file).Error; err == nil {
				file.App = *app
				return file, nil
			}
		}
	}

	var (
		err    error
		parent = &File{}
		parts  = strings.Split(strings.Trim(strings.TrimSpace(path), "/"), "/")
	)

	if parts[0] != "" {
		if parent, err = CreateOrGetRootPath(app, db); err != nil {
			return nil, err
		}
	}

	for _, part := range parts {
		var file = &File{}
		// deleted files should be considered
		if err = db.Where("appId = ? and pid = ? and name = ?", app.ID, parent.ID, part).First(file).Error; err != nil {
			return nil, err
		}
		parent = file
	}
	parent.App = *app

	_ = pathToFileCache.Add(cacheKey, parent, time.Minute*10)

	return parent, nil
}
