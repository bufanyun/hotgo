// Package tree
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tree

type Node interface {
	ID() int64                   // 获取节点ID
	PID() int64                  // 获取父级节点ID
	SetChildren(children []Node) // 设置子节点数据
}

// ListToTree 根据上下级关系将列表数据转为树状数据
func ListToTree(pid int64, nodes []Node) (list []Node, err error) {
	if len(nodes) == 0 {
		return nodes, nil
	}
	for _, v := range nodes {
		if v.PID() == pid {
			item := v

			// 递归添加子节点
			child, err := ListToTree(v.ID(), nodes)
			if err != nil {
				return nil, err
			}
			if len(child) > 0 {
				item.SetChildren(child)
			}
			list = append(list, item)
		}
	}
	return
}
