package main

import (
	"github.com/rustingoff/go-iris-api/controllers"
	"github.com/rustingoff/go-iris-api/interfaces"
	"github.com/rustingoff/go-iris-api/repositories"
	"github.com/rustingoff/go-iris-api/services"
	"sync"
)

type IServiceContainer interface {
	InjectTestController() interfaces.ITestController
}

type kernel struct {
}

func (k *kernel) InjectTestController() interfaces.ITestController {
	db := "sqlite3"
	testRepository := repositories.NewTestRepo(db)
	testService := services.NewTestService(testRepository)
	testController := controllers.NewTestController(testService)

	return testController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}

	return &kernel{}
}
