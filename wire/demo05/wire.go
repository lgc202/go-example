//go:build wireinject
// +build wireinject

package example

import "github.com/google/wire"

func GetPostService() (*PostService, func(), error) {
	panic(wire.Build(
		// 这里由于只有一个字段，所以这两种是等价的 wire.Struct(new(PostService), "*"),
		wire.Struct(new(PostService), "usecase"),
		wire.Bind(new(IPostUsecase), new(*PostUsecase)),
		NewPostUsecase,
		NewPostRepo,
	))
}
