// Package file
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package file

import (
	"github.com/gogf/gf/v2/os/gfile"
	"hotgo/utility/format"
	"os"
	"path/filepath"
)

// 文件信息
type fileInfo struct {
	name string
	size int64
}

// WalkDir 递归获取目录下文件的名称和大小
func WalkDir(dirname string) (error, []fileInfo) {
	op, err := filepath.Abs(dirname) // 获取目录的绝对路径
	if nil != err {
		return err, nil
	}
	files, err := os.ReadDir(op) // 获取目录下所有文件的信息，包括文件和文件夹
	if nil != err {
		return err, nil
	}

	var fileInfos []fileInfo // 返回值，存储读取的文件信息
	for _, f := range files {
		if f.IsDir() { // 如果是目录，那么就递归调用
			err, fs := WalkDir(op + `/` + f.Name()) // 路径分隔符，linux 和 windows 不同
			if nil != err {
				return err, nil
			}
			fileInfos = append(fileInfos, fs...) // 将 slice 添加到 slice
		} else {
			info, err := f.Info()
			if nil != err {
				return err, nil
			}
			fi := fileInfo{op + `/` + f.Name(), info.Size()}
			fileInfos = append(fileInfos, fi) // slice 中添加成员
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
