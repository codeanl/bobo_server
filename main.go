package main

import (
	"bobo_server/router"
	"golang.org/x/sync/errgroup"
	"log"
)

var g errgroup.Group

func main() {
	// 初始化全局变量
	router.InitAll()
	//
	// 前台接口服务
	g.Go(func() error {
		return router.FrontServer().ListenAndServe()
	})

	// 后台接口服务
	g.Go(func() error {
		return router.AdminServer().ListenAndServe()
	})
	//
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
