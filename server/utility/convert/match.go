package convert

import (
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/utility/validate"
)

// IpFilterStrategy ip过滤策略
func IpFilterStrategy(originIp string) (list map[string]struct{}) {
	list = make(map[string]struct{})

	// 所有IP
	if originIp == "*" {
		list["*"] = struct{}{}
		return
	}

	// 多个IP
	if gstr.Contains(originIp, ",") {
		ips := gstr.Explode(",", originIp)
		if len(ips) > 0 {
			for _, ip := range ips {
				if !validate.IsIp(ip) {
					continue
				}
				list[ip] = struct{}{}
			}
		}
		return
	}

	// IP段
	if gstr.Contains(originIp, "/24") {
		segment := gstr.Replace(originIp, "/24", "")
		if !validate.IsIp(segment) {
			return
		}

		var (
			start  = gstr.Explode(".", segment)
			prefix = gstr.Implode(".", start[:len(start)-1]) + "."
			index  = gconv.Int(start[len(start)-1])
		)

		if index < 1 {
			index = 1
		}

		for i := index; i <= 254; i++ {
			list[prefix+gconv.String(i)] = struct{}{}
		}
		return
	}

	// IP范围
	if gstr.Contains(originIp, "-") {
		originIps := gstr.Explode("-", originIp)
		if len(originIps) != 2 {
			return
		}

		if !validate.IsIp(originIps[0]) || !validate.IsIp(originIps[1]) {
			return
		}

		var (
			start      = gstr.Explode(".", originIps[0])
			prefix     = gstr.Implode(".", start[:len(start)-1]) + "."
			startIndex = gconv.Int(gstr.SubStrFromREx(originIps[0], "."))
			endIndex   = gconv.Int(gstr.SubStrFromREx(originIps[1], "."))
		)

		if startIndex >= endIndex {
			list[originIps[0]] = struct{}{}
			return
		}

		if startIndex < 1 {
			startIndex = 1
		}

		if endIndex > 254 {
			endIndex = 254
		}

		for i := startIndex; i <= endIndex; i++ {
			list[prefix+gconv.String(i)] = struct{}{}
		}
		return
	}

	// 指定IP
	if validate.IsIp(originIp) {
		list[originIp] = struct{}{}
		return
	}
	return list
}

// MatchIpStrategy 匹配IP策略，输入ip如果在策略当中返回true
func MatchIpStrategy(rules, ip string) bool {
	allowedIps := IpFilterStrategy(rules)
	if _, ok := allowedIps["*"]; ok {
		return true
	}

	_, ok := allowedIps[ip]
	return ok
}
