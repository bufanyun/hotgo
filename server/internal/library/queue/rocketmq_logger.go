// Package queue
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package queue

import (
	"fmt"
)

type RocketMqLogger struct {
	Flag     string
	LevelLog string
}

func (l *RocketMqLogger) Debug(msg string, fields map[string]interface{}) {
	if l.LevelLog == "close" {
		return
	}
	if msg == "" && len(fields) == 0 {
		return
	}

	if l.LevelLog == "debug" || l.LevelLog == "all" {
		Log(ctx, fmt.Sprint(l.Flag, " [debug] ", msg))
	}
}

func (l *RocketMqLogger) Level(level string) {
	Log(ctx, fmt.Sprint(l.Flag, " [level] ", level))
}

func (l *RocketMqLogger) OutputPath(path string) (err error) {
	Log(ctx, fmt.Sprint(l.Flag, " [path] ", path))
	return nil
}

func (l *RocketMqLogger) Info(msg string, fields map[string]interface{}) {
	if l.LevelLog == "close" {
		return
	}
	if msg == "" && len(fields) == 0 {
		return
	}

	if l.LevelLog == "info" || l.LevelLog == "all" {
		Log(ctx, fmt.Sprint(l.Flag, " [info] ", msg))
	}
}

func (l *RocketMqLogger) Warning(msg string, fields map[string]interface{}) {
	if l.LevelLog == "close" {
		return
	}
	if msg == "" && len(fields) == 0 {
		return
	}

	if l.LevelLog == "warn" || l.LevelLog == "all" {
		Log(ctx, fmt.Sprint(l.Flag, " [warn] ", msg))
	}
}

func (l *RocketMqLogger) Error(msg string, fields map[string]interface{}) {
	if l.LevelLog == "close" {
		return
	}
	if msg == "" && len(fields) == 0 {
		return
	}
	if l.LevelLog == "error" || l.LevelLog == "all" {
		Log(ctx, fmt.Sprint(l.Flag, " [error] ", msg))
	}
}

func (l *RocketMqLogger) Fatal(msg string, fields map[string]interface{}) {
	if l.LevelLog == "close" {
		return
	}
	if msg == "" && len(fields) == 0 {
		return
	}

	if l.LevelLog == "fatal" || l.LevelLog == "all" {
		Log(ctx, fmt.Sprint(l.Flag, " [fatal] ", msg))
	}
}
