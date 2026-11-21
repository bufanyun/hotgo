package queue

import (
	"errors"
	"strings"
	"testing"
)

func TestSafeReceiveDoRecoversPanic(t *testing.T) {
	err := safeReceiveDo(func(MqMsg) error {
		panic("boom")
	}, MqMsg{})
	if err == nil {
		t.Fatal("panic 应返回错误")
	}
	if !strings.Contains(err.Error(), "queue 消费回调 panic: boom") {
		t.Fatalf("错误信息不符: %v", err)
	}
}

func TestSafeReceiveDoPassesResult(t *testing.T) {
	want := errors.New("handle failed")
	if err := safeReceiveDo(func(MqMsg) error { return want }, MqMsg{}); !errors.Is(err, want) {
		t.Fatalf("应返回回调错误，得到 %v", err)
	}
	if err := safeReceiveDo(func(MqMsg) error { return nil }, MqMsg{}); err != nil {
		t.Fatalf("回调成功时应返回 nil，得到 %v", err)
	}
}
