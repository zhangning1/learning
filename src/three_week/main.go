package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	eg, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello Word ")
	})
	httpSvr := http.Server{Addr: "0.0.0.0:8088", Handler: mux}

	// 启动服务
	eg.Go(func() error {
		fmt.Println("启动服务")
		return httpSvr.ListenAndServe()
	})

	// 关闭服务
	eg.Go(func() error {
		stop := make(chan os.Signal)
		signal.Notify(stop, syscall.SIGINT, syscall.SIGUSR1)
		<-stop
		fmt.Println("关闭服务")
		return httpSvr.Shutdown(ctx)
	})

	// 监听服务
	eg.Go(func() error {
		ticker := time.NewTicker(time.Second * 10)
		for {
			select {
			case <-ticker.C:
				fmt.Println("监控程序 ： 服务正常运行中")
			case <-ctx.Done():
				fmt.Println("收到 ctx 信号, 退出监控程序")
				return nil
			}
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) &&
		!errors.Is(err, http.ErrServerClosed) {
		fmt.Println("有些错误哦，请确认 ：", err)
	}

	fmt.Println("程序退出")
}
