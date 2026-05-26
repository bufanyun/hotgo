// Package location
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package location

import (
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"sync"
)

type IpCache struct {
	sync.Mutex
	data *gmap.Map
}

var (
	cache = &IpCache{
		data: gmap.New(true),
	}
)

func (c *IpCache) Contains(ip string) bool {
	return c.data.Contains(ip)
}

func (c *IpCache) SetIpCache(ip string, data *IpLocationData) {
	if c.data.Size() > 10000 {
		c.data.Pops(2000)
	}
	c.data.Set(ip, data)
}

func (c *IpCache) GetIpCache(ip string) (*IpLocationData, error) {
	value := c.data.Get(ip)
	data1, ok := value.(*IpLocationData)
	if !ok {
		c.data.Remove(ip)
		err := fmt.Errorf("data assertion failed in the cache ip:%v", ip)
		return nil, err
	}
	return data1, nil
}
