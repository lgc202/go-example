//go:build wireinject
// +build wireinject

package example

import (
	"github.com/google/wire"
	"io"
	"os"
)

func GetPostService() (*PostService, func(), error) {
	panic(wire.Build(
		wire.Struct(new(PostService), "*"),
		wire.Value(10),
		wire.InterfaceValue(new(io.Reader), os.Stdin),
		wire.Bind(new(IPostUsecase), new(*PostUsecase)),
		NewPostUsecase,
		NewPostRepo,
	))
}
