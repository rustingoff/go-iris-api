package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/rustingoff/go-iris-api/interfaces"
)

type testController struct {
	interfaces.ITestService
}

func NewTestController(service interfaces.ITestService) *testController {
	return &testController{
		ITestService: service,
	}
}

func (controller *testController) Test(ctx iris.Context) {
	err := controller.ITestService.Test()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	options := iris.JSON{
		Indent: "  ",
		Secure: true,
	}

	ctx.StatusCode(iris.StatusCreated)
	err = ctx.JSON(iris.Map{
		"message": "ok",
	}, options)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}
}
