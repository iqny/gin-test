package main

import (
	"byobject/app/admin/main/user/conf"
	h "byobject/app/admin/main/user/http"
	"byobject/app/admin/main/user/service"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	s *service.Service
)

func main() {
	flag.Parse()
	c := conf.Init()
	s = service.New(c)
	defer s.Close()
	server := h.Init(c)
	//go server.ListenAndServeTLS("server.crt","server.key")
	go server.ListenAndServe()
	gracefulExitWeb(server)
}
func gracefulExitWeb(server *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch

	fmt.Println("got a signal", sig)
	now := time.Now()
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(cxt)
	if err != nil {
		fmt.Println("err", err)
	}
	// 看看实际退出所耗费的时间
	fmt.Println("------exited--------", time.Since(now))
}
