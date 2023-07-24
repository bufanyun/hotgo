// Package websocket
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package websocket

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
	"runtime/debug"
	"sync"
)

// ClientManager 客户端管理
type ClientManager struct {
	Clients         map[*Client]bool      // 全部的连接
	ClientsLock     sync.RWMutex          // 读写锁
	Users           map[string][]*Client  // 登录的用户
	UserLock        sync.RWMutex          // 读写锁
	Register        chan *Client          // 连接连接处理
	Login           chan *login           // 用户登录处理
	Unregister      chan *Client          // 断开连接处理程序
	Broadcast       chan *WResponse       // 广播 向全部成员发送数据
	ClientBroadcast chan *ClientWResponse // 广播 向某个客户端发送数据
	TagBroadcast    chan *TagWResponse    // 广播 向某个标签成员发送数据
	UserBroadcast   chan *UserWResponse   // 广播 向某个用户的所有链接发送数据
	closeSignal     chan struct{}         // 关闭信号
}

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		Clients:       make(map[*Client]bool),
		Users:         make(map[string][]*Client),
		Register:      make(chan *Client, 1000),
		Unregister:    make(chan *Client, 1000),
		Broadcast:     make(chan *WResponse, 1000),
		TagBroadcast:  make(chan *TagWResponse, 1000),
		UserBroadcast: make(chan *UserWResponse, 1000),
		closeSignal:   make(chan struct{}, 1),
	}
	return
}

func Manager() *ClientManager {
	return clientManager
}

// GetUserKey 获取用户key
func GetUserKey(userId int64) (key string) {
	key = fmt.Sprintf("%s_%d", "ws", userId)
	return
}

// InClient 客户端是否存在
func (manager *ClientManager) InClient(client *Client) (ok bool) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()
	_, ok = manager.Clients[client]
	return
}

// GetClients 获取所有客户端
func (manager *ClientManager) GetClients() (clients map[*Client]bool) {
	clients = make(map[*Client]bool)
	manager.ClientsRange(func(client *Client, value bool) (result bool) {
		clients[client] = value
		return true
	})
	return
}

// ClientsRange 遍历
func (manager *ClientManager) ClientsRange(f func(client *Client, value bool) (result bool)) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()
	for key, value := range manager.Clients {
		if !f(key, value) {
			return
		}
	}
}

// GetClientsLen 获取客户端总数
func (manager *ClientManager) GetClientsLen() (clientsLen int) {
	clientsLen = len(manager.Clients)
	return
}

// AddClients 添加客户端
func (manager *ClientManager) AddClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()
	manager.Clients[client] = true
}

// DelClients 删除客户端
func (manager *ClientManager) DelClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()
	delete(manager.Clients, client)
}

// GetClient 通过socket ID获取客户端的连接
func (manager *ClientManager) GetClient(id string) (client *Client) {
	for c := range manager.Clients {
		if c.ID == id {
			return c
		}
	}
	return
}

// GetUserClient 获取用户的连接
func (manager *ClientManager) GetUserClient(userId int64) (clients []*Client) {
	manager.UserLock.RLock()
	defer manager.UserLock.RUnlock()
	userKey := GetUserKey(userId)
	if value, ok := manager.Users[userKey]; ok {
		clients = value
	}
	return
}

// AddUsers 添加用户
func (manager *ClientManager) AddUsers(key string, client *Client) {
	manager.UserLock.Lock()
	defer manager.UserLock.Unlock()
	manager.Users[key] = append(manager.Users[key], client)
}

// DelUsers 删除用户
func (manager *ClientManager) DelUsers(client *Client) (result bool) {
	manager.UserLock.Lock()
	defer manager.UserLock.Unlock()
	key := GetUserKey(client.User.Id)
	if clients, ok := manager.Users[key]; ok {
		for _, value := range clients {
			// 判断是否为相同的用户
			if value.Addr != client.Addr {
				return
			}
			delete(manager.Users, key)
			result = true
		}
	}
	return
}

// GetUsersLen 已登录用户数
func (manager *ClientManager) GetUsersLen() (userLen int) {
	userLen = len(manager.Users)
	return
}

// EventRegister 用户建立连接事件
func (manager *ClientManager) EventRegister(client *Client) {
	manager.AddClients(client)
	// 用户登录
	manager.EventLogin(&login{
		UserId: client.User.Id,
		Client: client,
	})
	// 发送当前客户端标识
	// SendSuccess(client, "connected", g.Map{"id": client.ID, "userInfo": client.User})
}

// EventLogin 用户登录事件
func (manager *ClientManager) EventLogin(login *login) {
	client := login.Client
	if manager.InClient(client) {
		userKey := login.GetKey()
		manager.AddUsers(userKey, login.Client)
	}
}

// EventUnregister 用户断开连接事件
func (manager *ClientManager) EventUnregister(client *Client) {
	manager.DelClients(client)
	// 删除用户连接
	deleteResult := manager.DelUsers(client)
	if !deleteResult {
		// 不是当前连接的客户端
		return
	}

	client.close()
}

// ClearTimeoutConnections 定时清理超时连接
func (manager *ClientManager) clearTimeoutConnections() {
	currentTime := uint64(gtime.Now().Unix())
	clients := clientManager.GetClients()
	for client := range clients {
		if client.IsHeartbeatTimeout(currentTime) {
			_ = client.Socket.Close()
		}
	}
}

// WebsocketPing 心跳处理
func (manager *ClientManager) ping() {
	defer func() {
		if r := recover(); r != nil {
			g.Log().Warningf(ctxManager, "websocket gcron ping recover:%+v, stack:%+v", r, string(debug.Stack()))
			return
		}
	}()

	// 定时任务，发送心跳包
	// gcron.Add(ctx, "0 */1 * * * *", func(ctx context.Context) {
	//	res := &WResponse{
	//		Event:     "ping",
	// 		Timestamp: gtime.Now().Unix(),
	// 	}
	// 	SendToAll(res)
	// })

	// 定时任务，清理超时连接
	_, _ = gcron.Add(ctxManager, "*/30 * * * * *", func(ctx context.Context) {
		manager.clearTimeoutConnections()
	})
}

// 管道处理程序
func (manager *ClientManager) start() {
	defer func() {
		if r := recover(); r != nil {
			g.Log().Warningf(ctxManager, "websocket start recover:%+v, stack:%+v", r, string(debug.Stack()))
			return
		}
	}()
	for {
		select {
		case conn := <-manager.Register:
			// 建立连接事件
			manager.EventRegister(conn)
		case login := <-manager.Login:
			// 用户登录
			manager.EventLogin(login)

		case conn := <-manager.Unregister:
			// 断开连接事件
			manager.EventUnregister(conn)

		case message := <-manager.Broadcast:
			// 全部客户端广播事件
			clients := manager.GetClients()
			for conn := range clients {
				conn.SendMsg(message)
			}
		case message := <-manager.TagBroadcast:
			// 标签广播事件
			clients := manager.GetClients()
			for conn := range clients {
				if conn.Tags.Contains(message.Tag) {
					if message.WResponse.Timestamp == 0 {
						message.WResponse.Timestamp = gtime.Now().Timestamp()
					}
					conn.SendMsg(message.WResponse)
				}
			}
		case message := <-manager.UserBroadcast:
			// 用户广播事件
			clients := manager.GetClients()

			for conn := range clients {
				if conn.User.Id == message.UserID {
					if message.WResponse.Timestamp == 0 {
						message.WResponse.Timestamp = gtime.Now().Timestamp()
					}
					conn.SendMsg(message.WResponse)
				}
			}
		case message := <-manager.ClientBroadcast:
			// 单个客户端广播事件
			clients := manager.GetClients()
			for conn := range clients {
				if conn.ID == message.ID {
					if message.WResponse.Timestamp == 0 {
						message.WResponse.Timestamp = gtime.Now().Timestamp()
					}
					conn.SendMsg(message.WResponse)
				}
			}
		case <-manager.closeSignal:
			g.Log().Debug(ctxManager, "websocket closeSignal quit..")
			return
		}
	}
}

// SendToAll 发送全部客户端
func SendToAll(response *WResponse) {
	clientManager.Broadcast <- response
}

// SendToClientID  发送单个客户端
func SendToClientID(id string, response *WResponse) {
	clientRes := &ClientWResponse{
		ID:        id,
		WResponse: response,
	}
	clientManager.ClientBroadcast <- clientRes
}

// SendToUser 发送单个用户
func SendToUser(userID int64, response *WResponse) {
	userRes := &UserWResponse{
		UserID:    userID,
		WResponse: response,
	}
	clientManager.UserBroadcast <- userRes
}

// SendToTag 发送某个标签
func SendToTag(tag string, response *WResponse) {
	tagRes := &TagWResponse{
		Tag:       tag,
		WResponse: response,
	}
	clientManager.TagBroadcast <- tagRes
}
