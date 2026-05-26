// Package dict
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package dict

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"hash/fnv"
	"hotgo/internal/model"
	"strconv"
)

const (
	BuiltinId int64 = -1 // 内置字典ID
	EnumsId   int64 = -2 // 枚举字典ID
	FuncId    int64 = -3 // 方法字典ID
)

var NotExistKeyError = errors.New("not exist key")

// GetOptions 获取内置选项
func GetOptions(ctx context.Context, key string) (opts []*model.Option, err error) {
	opts = GetEnumsOptions(key)
	if opts != nil {
		return
	}
	return GetFuncOptions(ctx, key)
}

// GetOptionsById 通过类型ID获取内置选项
func GetOptionsById(ctx context.Context, id int64) (opts []*model.Option, err error) {
	for _, v := range GetAllEnums() {
		if v.Id == id {
			return v.Opts, nil
		}
	}

	for _, v := range GetAllFunc() {
		g.Log().Warningf(ctx, "GetAllFunc GetOptionsById v:%v, %v", v.Id, v.Key)
		if v.Id == id {
			return LoadFuncOptions(ctx, v)
		}
	}

	err = NotExistKeyError
	return
}

// GetTypeById 通过类型ID获取内置选项类型
func GetTypeById(ctx context.Context, id int64) (typ string, err error) {
	for _, v := range GetAllEnums() {
		if v.Id == id {
			return v.Key, nil
		}
	}

	for _, v := range GetAllFunc() {
		if v.Id == id {
			return v.Key, nil
		}
	}

	err = NotExistKeyError
	return
}

// GenIdHash 生成字典id
func GenIdHash(str string, t int64) int64 {
	prefix := 10000 * t
	h := fnv.New32a()
	h.Write([]byte("dict" + str))

	idStr := fmt.Sprintf("%d%d", prefix, int64(h.Sum32()))
	id, _ := strconv.ParseInt(idStr, 10, 64)
	return id
}
