//go:build wireinject
// +build wireinject

package example

import "github.com/google/wire"

func GetPostService() (*PostService, error) {
	panic(wire.Build(
		NewPostService,
		NewPostUsecase,
		NewPostRepo,
	))
}
