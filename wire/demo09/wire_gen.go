// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package example

import (
	"os"
)

// Injectors from wire.go:

func GetPostService() (*PostService, func(), error) {
	a := _wireAValue
	b := _wireBValue
	iPostRepo, cleanup, err := NewPostRepo()
	if err != nil {
		return nil, nil, err
	}
	postUsecaseOption := &PostUsecaseOption{
		a:    a,
		b:    b,
		repo: iPostRepo,
	}
	postUsecase, cleanup2, err := NewPostUsecase(postUsecaseOption)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	reader := _wireFileValue
	postService := &PostService{
		usecase: postUsecase,
		a:       a,
		b:       b,
		r:       reader,
	}
	return postService, func() {
		cleanup2()
		cleanup()
	}, nil
}

var (
	_wireAValue    = A(10)
	_wireBValue    = B(10)
	_wireFileValue = os.Stdin
)