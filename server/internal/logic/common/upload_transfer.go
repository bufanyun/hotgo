package common

import (
	"io"
	"net"
	"net/http"
	"net/netip"
	"syscall"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/internal/library/storager"
)

const defaultImageSizeMB int64 = 10

var (
	errPrivateAddr      = gerror.New("不允许访问内网地址")
	errTooManyRedirects = gerror.New("重定向次数过多")
	errBadScheme        = gerror.New("仅支持 HTTP/HTTPS 协议的图片链接")
	errImageTooLarge    = gerror.New("图片大小超过限制")
)

// 这些网段不能作为转存目标，防止探测内网
var nonPublicPrefixes = []netip.Prefix{
	netip.MustParsePrefix("0.0.0.0/8"),
	netip.MustParsePrefix("10.0.0.0/8"),
	netip.MustParsePrefix("100.64.0.0/10"),
	netip.MustParsePrefix("127.0.0.0/8"),
	netip.MustParsePrefix("169.254.0.0/16"),
	netip.MustParsePrefix("172.16.0.0/12"),
	netip.MustParsePrefix("192.0.0.0/24"),
	netip.MustParsePrefix("192.0.2.0/24"),
	netip.MustParsePrefix("192.168.0.0/16"),
	netip.MustParsePrefix("198.18.0.0/15"),
	netip.MustParsePrefix("198.51.100.0/24"),
	netip.MustParsePrefix("203.0.113.0/24"),
	netip.MustParsePrefix("224.0.0.0/4"),
	netip.MustParsePrefix("240.0.0.0/4"),
	netip.MustParsePrefix("::/128"),
	netip.MustParsePrefix("::1/128"),
	netip.MustParsePrefix("::/96"),
	netip.MustParsePrefix("64:ff9b::/96"),
	netip.MustParsePrefix("64:ff9b:1::/48"),
	netip.MustParsePrefix("100::/64"),
	netip.MustParsePrefix("2001:db8::/32"),
	netip.MustParsePrefix("fc00::/7"),
	netip.MustParsePrefix("fe80::/10"),
	netip.MustParsePrefix("fec0::/10"),
	netip.MustParsePrefix("ff00::/8"),
}

func isPublicIP(ip net.IP) bool {
	addr, ok := netip.AddrFromSlice(ip)
	if !ok {
		return false
	}
	// IPv4 映射地址要按 IPv4 判断，否则绕过 127.0.0.0/8 等前缀
	addr = addr.Unmap()
	for _, prefix := range nonPublicPrefixes {
		if prefix.Contains(addr) {
			return false
		}
	}
	return true
}

func newTransferClient(allowPrivate bool) *http.Client {
	dialer := &net.Dialer{
		Timeout: 10 * time.Second,
		Control: func(network, address string, _ syscall.RawConn) error {
			if allowPrivate {
				return nil
			}
			host, _, err := net.SplitHostPort(address)
			if err != nil {
				return gerror.Wrap(err, "解析目标地址失败")
			}
			ip := net.ParseIP(host)
			if ip == nil || !isPublicIP(ip) {
				return errPrivateAddr
			}
			return nil
		},
	}

	transport := &http.Transport{
		// 走代理时 Control 只能看到代理地址
		Proxy:       nil,
		DialContext: dialer.DialContext,
		// 每次请求新建 client，留空闲连接会泄漏连接与协程
		DisableKeepAlives:     true,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 15 * time.Second,
		MaxIdleConns:          0,
	}

	return &http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if req.URL.Scheme != "http" && req.URL.Scheme != "https" {
				return errBadScheme
			}
			if len(via) > 3 {
				return errTooManyRedirects
			}
			return nil
		},
	}
}

func imageSizeLimitMB() int64 {
	if cfg := storager.GetConfig(); cfg != nil && cfg.ImageSize > 0 {
		return cfg.ImageSize
	}
	return defaultImageSizeMB
}

func readImageBody(body io.Reader, limitMB int64) ([]byte, error) {
	max := limitMB * 1024 * 1024
	content, err := io.ReadAll(io.LimitReader(body, max+1))
	if err != nil {
		return nil, err
	}
	if int64(len(content)) > max {
		return nil, gerror.Wrapf(errImageTooLarge, "图片大小不能超过%vMB", limitMB)
	}
	return content, nil
}
