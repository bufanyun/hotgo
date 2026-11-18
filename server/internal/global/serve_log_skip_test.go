package global

import (
	"context"
	"testing"
)

func TestIsServeLogWrite(t *testing.T) {
	ctx := context.Background()
	if IsServeLogWrite(ctx) {
		t.Fatal("普通上下文不应带写服务日志标记")
	}

	marked := MarkServeLogWrite(ctx)
	if !IsServeLogWrite(marked) {
		t.Fatal("标记后的上下文丢失了写服务日志标记")
	}

	child := context.WithValue(marked, struct{ name string }{"child"}, 1)
	if !IsServeLogWrite(child) {
		t.Fatal("子上下文丢失了写服务日志标记")
	}

	other := context.WithValue(ctx, struct{ name string }{"other"}, 1)
	if IsServeLogWrite(other) {
		t.Fatal("无关上下文不应命中写服务日志标记")
	}
}
