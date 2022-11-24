// Package signal
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package signal

import (
	"sync"
)

type StopSignal int32

type exitWait struct {
	mutex        sync.Mutex
	wg           *sync.WaitGroup
	deferFuns    []func()
	stopSignList []chan StopSignal
}

var exitWaitHandler *exitWait

func init() {
	exitWaitHandler = &exitWait{
		wg: &sync.WaitGroup{},
	}
}

// ExitWaitFunDo 退出后等待处理完成
func ExitWaitFunDo(doFun func()) {
	exitWaitHandler.wg.Add(1)
	defer exitWaitHandler.wg.Done()
	if doFun != nil {
		doFun()
	}
}

// AppDefer 应用退出后置操作
func AppDefer(deferFun ...func()) {
	exitWaitHandler.mutex.Lock()
	defer exitWaitHandler.mutex.Unlock()
	for _, funcItem := range deferFun {
		if funcItem != nil {
			exitWaitHandler.deferFuns = append(exitWaitHandler.deferFuns, funcItem)
		}
	}
}

// ListenStop 订阅app退出信号
func ListenStop(stopSig chan StopSignal) {
	exitWaitHandler.mutex.Lock()
	defer exitWaitHandler.mutex.Unlock()

	exitWaitHandler.stopSignList = append(exitWaitHandler.stopSignList, stopSig)
}
