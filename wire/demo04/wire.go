//go:build wireinject
// +build wireinject

package example

import "github.com/google/wire"

func GetPostService() (*PostService, func(), error) {
	panic(wire.Build(
		NewPostService,
		wire.Bind(new(IPostUsecase), new(*PostUsecase)),
		NewPostUsecase,
		NewPostRepo,
	))
}
