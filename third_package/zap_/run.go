package main

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

func main() {
	// 创建一个新的 zap 日志器
	logger, _ := zap.NewProduction()

	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			// 这里实际上可以忽略这个错误，因为日志库已经将错误写入到日志中
			// 但是为了严谨，还是打印出来
			fmt.Println("日志同步失败", err)
			// panic(err)
		}
	}(logger)
	// 记录一条信息级别的日志
	logger.Info("这是一条信息级别的日志",
		// 结构化字段
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", 2*time.Second),
	)

	// 记录一条错误级别的日志
	logger.Error("这是一条错误级别的日志",
		// 结构化字段
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
