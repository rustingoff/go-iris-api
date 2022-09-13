package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/rustingoff/go-iris-api/packages/logger"
	"sync"
)

type IIrisRouter interface {
	InitRouter() *iris.Application
}

type router struct {
}

func (r *router) InitRouter() *iris.Application {
	accessLog := logger.MakeAccessLog()
	defer func() {
		err := accessLog.Close()
		if err != nil {
			panic(err)
		}
	}()

	app := iris.New()
	app.Use(accessLog.Handler)
	app.UseRouter(recover.New())

	testController := ServiceContainer().InjectTestController()

	testRoute := app.Party("/test")
	{
		testRoute.Get("/", testController.Test)
	}

	err := app.Build()
	if err != nil {
		app.Logger().Fatal(err)
	}

	return app
}

var (
	m          *router
	routerOnce sync.Once
)

func IrisRouter() IIrisRouter {
	routerOnce.Do(func() {
		m = &router{}
	})

	return m
}
