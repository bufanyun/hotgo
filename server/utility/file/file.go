// Package file
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package file

import (
	"github.com/gogf/gf/v2/os/gfile"
	"hotgo/utility/format"
	"io/ioutil"
	"os"
	"path/filepath"
)

const ( //文件大小单位
	_  = iota
	KB = 1 << (10 * iota)
	MB
)

type fileInfo struct { //文件信息
	name string
	size int64
}

func PathExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir(), nil
	}

	return false, err
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// HasDir 判断文件夹是否存在
func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

// CreateDir 创建文件夹
func CreateDir(path string) (err error) {
	_exist, err := HasDir(path)
	if err != nil {
		return
	}
	if !_exist {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return
		}
	}
	return
}

// WalkDir 递归获取目录下文件的名称和大小
func WalkDir(dirname string) (error, []fileInfo) {
	op, err := filepath.Abs(dirname) //获取目录的绝对路径
	if nil != err {
		return err, nil
	}
	files, err := ioutil.ReadDir(op) //获取目录下所有文件的信息，包括文件和文件夹
	if nil != err {
		return err, nil
	}

	var fileInfos []fileInfo //返回值，存储读取的文件信息
	for _, f := range files {
		if f.IsDir() { // 如果是目录，那么就递归调用
			err, fs := WalkDir(op + `/` + f.Name()) //路径分隔符，linux 和 windows 不同
			if nil != err {
				return err, nil
			}
			fileInfos = append(fileInfos, fs...) //将 slice 添加到 slice
		} else {
			fi := fileInfo{op + `/` + f.Name(), f.Size()}
			fileInfos = append(fileInfos, fi) //slice 中添加成员
		}
	}

	return nil, fileInfos
}

// DirSize 获取目录下所有文件大小
func DirSize(dirname string) string {
	var (
		ss       int64
		_, files = WalkDir(dirname)
	)
	for _, n := range files {
		ss += n.size
	}

	return format.FileSize(ss)
}

func MergeAbs(path string, fileName ...string) string {
	var paths = []string{gfile.RealPath(path)}
	paths = append(paths, fileName...)
	return gfile.Join(paths...)
}
