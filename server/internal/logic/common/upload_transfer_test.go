package common

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestIsPublicIP(t *testing.T) {
	cases := []struct {
		ip     string
		public bool
	}{
		{"0.0.0.1", false},
		{"10.0.0.1", false},
		{"100.64.0.1", false},
		{"127.0.0.1", false},
		{"169.254.169.254", false},
		{"172.16.0.1", false},
		{"192.0.0.1", false},
		{"192.0.2.1", false},
		{"192.168.1.1", false},
		{"198.18.0.1", false},
		{"198.51.100.1", false},
		{"203.0.113.1", false},
		{"224.0.0.1", false},
		{"240.0.0.1", false},
		{"::", false},
		{"::1", false},
		{"::7f00:1", false},
		{"64:ff9b::7f00:1", false},
		{"64:ff9b:1::1", false},
		{"100::1", false},
		{"2001:db8::1", false},
		{"fc00::1", false},
		{"fe80::1", false},
		{"fec0::1", false},
		{"ff02::1", false},
		{"::ffff:127.0.0.1", false},
		{"8.8.8.8", true},
		{"1.1.1.1", true},
		{"2001:4860:4860::8888", true},
	}
	for _, c := range cases {
		if got := isPublicIP(net.ParseIP(c.ip)); got != c.public {
			t.Errorf("isPublicIP(%s) = %v, want %v", c.ip, got, c.public)
		}
	}
}

func newImageServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		_, _ = w.Write([]byte("png"))
	}))
}

func TestTransferClientRejectsPrivateAddr(t *testing.T) {
	server := newImageServer()
	defer server.Close()

	_, err := newTransferClient(false).Get(server.URL)
	if !errors.Is(err, errPrivateAddr) {
		t.Fatalf("newTransferClient(false) err = %v, want errPrivateAddr", err)
	}

	resp, err := newTransferClient(true).Get(server.URL)
	if err != nil {
		t.Fatalf("newTransferClient(true) err = %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil || string(body) != "png" {
		t.Fatalf("newTransferClient(true) body = %q, err = %v", body, err)
	}
}

func TestTransferClientProxyNil(t *testing.T) {
	transport, ok := newTransferClient(false).Transport.(*http.Transport)
	if !ok {
		t.Fatalf("Transport 类型 = %T, want *http.Transport", newTransferClient(false).Transport)
	}
	if transport.Proxy != nil {
		t.Fatalf("newTransferClient(false) Proxy 不为 nil")
	}
}

func TestTransferClientRejectsBadRedirectScheme(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "ftp://x/image.png", http.StatusFound)
	}))
	defer server.Close()

	_, err := newTransferClient(true).Get(server.URL)
	if !errors.Is(err, errBadScheme) {
		t.Fatalf("重定向到 ftp err = %v, want errBadScheme", err)
	}
}

func TestTransferClientRedirectLimit(t *testing.T) {
	newRedirectServer := func(redirects int) *httptest.Server {
		return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hop, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/"))
			if hop < redirects {
				http.Redirect(w, r, fmt.Sprintf("/%d", hop+1), http.StatusFound)
				return
			}
			_, _ = w.Write([]byte("ok"))
		}))
	}

	allowed := newRedirectServer(3)
	defer allowed.Close()
	resp, err := newTransferClient(true).Get(allowed.URL + "/0")
	if err != nil {
		t.Fatalf("3 次重定向 err = %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil || string(body) != "ok" {
		t.Fatalf("3 次重定向 body = %q, err = %v", body, err)
	}

	tooMany := newRedirectServer(4)
	defer tooMany.Close()
	_, err = newTransferClient(true).Get(tooMany.URL + "/0")
	if !errors.Is(err, errTooManyRedirects) {
		t.Fatalf("4 次重定向 err = %v, want errTooManyRedirects", err)
	}
}

func TestReadImageBodyLimit(t *testing.T) {
	const limitMB int64 = 1
	size := int(limitMB * 1024 * 1024)

	content, err := readImageBody(bytes.NewReader(bytes.Repeat([]byte{'x'}, size)), limitMB)
	if err != nil || len(content) != size {
		t.Fatalf("readImageBody(limit) len = %d, err = %v", len(content), err)
	}

	_, err = readImageBody(bytes.NewReader(bytes.Repeat([]byte{'x'}, size+1)), limitMB)
	if !errors.Is(err, errImageTooLarge) {
		t.Fatalf("readImageBody(over limit) err = %v, want errImageTooLarge", err)
	}
	if !strings.Contains(err.Error(), "图片大小不能超过1MB") {
		t.Fatalf("readImageBody(over limit) message = %v", err)
	}
}
