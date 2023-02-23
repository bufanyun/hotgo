// Package encrypt
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package encrypt

import (
	"crypto/md5"
	"fmt"
	"hash/fnv"
)

// Md5ToString 生成md5
func Md5ToString(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// Md5 生成md5
func Md5(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

func Hash32(b []byte) uint32 {
	h := fnv.New32a()
	h.Write(b)
	return h.Sum32()
}
