// Package casbin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package casbin

import (
	"context"
	"net/http"
	"testing"
)

// TestNew description
func TestNew(t *testing.T) {

	InitEnforcer(context.TODO())

	user := "admin"
	path := "/"
	method := http.MethodGet
	t.Logf("\nuser:%v\npath:%v\nmethod:%v", user, path, method)

	ok, err := Enforcer.DeletePermissionsForUser(user)
	if err != nil {
		t.Error(err)
	}
	t.Logf("delete user premission:%v", ok)
	CheckPremission(t, user, path, method)
	AddPremission(t, user, "*", ActionAll)
	CheckPremission(t, user, path, method)

	user1 := "user1"
	path1 := "/api/v1/*"
	checkPathTrue := "/api/v1/user/list"
	checkPathFalse := "/api/v2/user/list"
	AddPremission(t, user1, path1, ActionAll)
	CheckPremission(t, user1, checkPathTrue, ActionPost)
	CheckPremission(t, user1, checkPathFalse, http.MethodGet)
	CheckPremission(t, user1, checkPathTrue, http.MethodGet)
	CheckPremission(t, user1, "/api/v1/user/list2", http.MethodGet)

	ok, err = Enforcer.DeletePermissionsForUser(user1)
	if err != nil {
		t.Error(err)
	}
	t.Logf("delete user premission:%v", ok)
	CheckPremission(t, user1, "/api/v1/user1/list", http.MethodGet)
}

// CheckPremission description
func CheckPremission(t *testing.T, user string, path string, method string) {
	ok, err := Enforcer.Enforce(user, path, method)
	if err != nil {
		t.Error(err)
	}
	t.Logf("check \tuser[%s] \tpremission[%s] \tpath[%s] \tallow[%v]", user, method, path, ok)
}

// Add description
func AddPremission(t *testing.T, user string, path string, method string) {
	ok, err := Enforcer.AddPolicy(user, path, method)
	if err != nil {
		t.Error(err)
	}
	t.Logf("add \tuser[%s] \tpremission[%s] \tpath[%s] \tresult[%v]", user, method, path, ok)
}
