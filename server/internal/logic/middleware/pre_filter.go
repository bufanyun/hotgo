// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package middleware

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/library/response"
	"hotgo/utility/validate"
	"reflect"
)

// GetFilterRoutes 获取支持预处理的web路由
func (s *sMiddleware) GetFilterRoutes(r *ghttp.Request) map[string]ghttp.RouterItem {
	// 首次访问时加载
	if s.FilterRoutes == nil {
		s.routeMutex.Lock()
		defer s.routeMutex.Unlock()

		if s.FilterRoutes != nil {
			return s.FilterRoutes
		}

		s.FilterRoutes = make(map[string]ghttp.RouterItem)
		for _, v := range r.Server.GetRoutes() {
			// 非规范路由不加载
			if v.Handler.Info.Type.NumIn() != 2 {
				continue
			}

			key := s.GenFilterRouteKey(v.Handler.Router)
			if _, ok := s.FilterRoutes[key]; !ok {
				s.FilterRoutes[key] = v
			}
		}
	}
	return s.FilterRoutes
}

// GenFilterRequestKey 根据请求生成唯一key
func (s *sMiddleware) GenFilterRequestKey(r *ghttp.Request) string {
	return s.GenRouteKey(r.Method, r.Request.URL.Path)
}

// GenFilterRouteKey 根据路由生成唯一key
func (s *sMiddleware) GenFilterRouteKey(r *ghttp.Router) string {
	return s.GenRouteKey(r.Method, r.Uri)
}

// GenRouteKey 生成唯一key
func (s *sMiddleware) GenRouteKey(method, path string) string {
	return method + " " + path
}

// PreFilter 请求输入预处理
// api使用gf规范路由并且XxxReq结构体实现了validate.Filter接口即可
func (s *sMiddleware) PreFilter(r *ghttp.Request) {
	router, ok := s.GetFilterRoutes(r)[s.GenFilterRequestKey(r)]
	if !ok {
		r.Middleware.Next()
		return
	}

	funcInfo := router.Handler.Info

	// 非规范路由不处理
	if funcInfo.Type.NumIn() != 2 {
		r.Middleware.Next()
		return
	}

	inputType := funcInfo.Type.In(1)
	var inputObject reflect.Value
	if inputType.Kind() == reflect.Ptr {
		inputObject = reflect.New(inputType.Elem())
	} else {
		inputObject = reflect.New(inputType.Elem()).Elem()
	}

	// 先验证基本校验规则
	if err := r.Parse(inputObject.Interface()); err != nil {
		resp := gerror.Code(err)
		response.JsonExit(r, resp.Code(), gerror.Current(err).Error(), resp.Detail())
		return
	}

	// 没有实现预处理
	if _, ok = inputObject.Interface().(validate.Filter); !ok {
		r.Middleware.Next()
		return
	}

	// 执行预处理
	if err := validate.PreFilter(r.Context(), inputObject.Interface()); err != nil {
		resp := gerror.Code(err)
		response.JsonExit(r, resp.Code(), gerror.Current(err).Error(), resp.Detail())
		return
	}

	// 过滤后的参数
	r.SetParamMap(gconv.Map(inputObject.Interface()))
	r.Middleware.Next()
}
