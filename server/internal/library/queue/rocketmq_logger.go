// Package queue
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package queue

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
		Logger().Debug(ctx, msg)
	}
}

func (l *RocketMqLogger) Level(level string) {
	Logger().Info(ctx, level)
}

func (l *RocketMqLogger) OutputPath(path string) (err error) {
	Logger().Info(ctx, path)
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
		Logger().Info(ctx, msg)
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
		Logger().Warning(ctx, msg)
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
		Logger().Error(ctx, msg)
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
		Logger().Fatal(ctx, msg)
	}
}
