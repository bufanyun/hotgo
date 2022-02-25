//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package utils

import (
	"sync"
)

// 信号类
var Signal = new(signal)

type signal struct{}

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

//
//  @Title  退出后等待处理完成
//  @Description 
//  @Author  Ms <133814250@qq.com>
//  @Param   doFun
//
func (util *signal) ExitWaitFunDo(doFun func()) {
	exitWaitHandler.wg.Add(1)
	defer exitWaitHandler.wg.Done()
	if doFun != nil {
		doFun()
	}
}

//
//  @Title  应用退出后置操作
//  @Description 
//  @Author  Ms <133814250@qq.com>
//  @Param   deferFun
//
func (util *signal) AppDefer(deferFun ...func()) {
	exitWaitHandler.mutex.Lock()
	defer exitWaitHandler.mutex.Unlock()
	for _, funcItem := range deferFun {
		if funcItem != nil {
			exitWaitHandler.deferFuns = append(exitWaitHandler.deferFuns, funcItem)
		}
	}
}

//
//  @Title  订阅app退出信号
//  @Description 
//  @Author  Ms <133814250@qq.com>
//  @Param   stopSig
//
func (util *signal) ListenStop(stopSig chan StopSignal) {
	exitWaitHandler.mutex.Lock()
	defer exitWaitHandler.mutex.Unlock()

	exitWaitHandler.stopSignList = append(exitWaitHandler.stopSignList, stopSig)
}
