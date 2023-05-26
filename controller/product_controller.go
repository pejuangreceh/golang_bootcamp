package controller

import "golang_project/model"

func Create(Id uint) {
	model.Create(Id)
}

type Controller struct {
	useCase UseCase
}

type UseCase interface {
}

func New() Controller {
	return Controller{}
}

type useCase struct {
}

var defaultController *Controller

// func Default() Controller {
//
//		return Controller{
//			useCase: useCase{},
//		}
//	}
func Default() *Controller {
	if defaultController == nil {
		defaultController = &Controller{useCase: useCase{}}

	}
	return defaultController
}
