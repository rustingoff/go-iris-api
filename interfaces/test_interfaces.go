package interfaces

import "github.com/kataras/iris/v12"

type ITestRepo interface {
	Test() (string, error)
}

type ITestService interface {
	Test() error
}

type ITestController interface {
	Test(ctx iris.Context)
}
