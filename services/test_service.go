package services

import "github.com/rustingoff/go-iris-api/interfaces"

type testService struct {
	interfaces.ITestRepo
}

func NewTestService(repo interfaces.ITestRepo) interfaces.ITestService {
	return &testService{
		ITestRepo: repo,
	}
}

func (service *testService) Test() error {
	_, err := service.ITestRepo.Test()
	return err
}
