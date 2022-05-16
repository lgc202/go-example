package example

import (
	"github.com/google/wire"
	"io"
	"os"
)

// repo

// IPostRepo IPostRepo
type IPostRepo interface{}

// NewPostRepo NewPostRepo
func NewPostRepo() (IPostRepo, func(), error) {
	return new(IPostRepo), nil, nil
}

// usecase

// IPostUsecase IPostUsecase
type IPostUsecase interface{}
type PostUsecase struct {
	repo IPostRepo
}

// NewPostUsecase NewPostUsecase
func NewPostUsecase(repo IPostRepo) (*PostUsecase, func(), error) {
	return &PostUsecase{repo: repo}, nil, nil
}

// service service

// PostService PostService
type PostService struct {
	usecase IPostUsecase
	a       int
	r       io.Reader
}

// NewPostService NewPostService
func NewPostService(u IPostUsecase) (*PostService, error) {
	return &PostService{usecase: u}, nil
}

// PostServiceSet PostServiceSet
var PostServiceSet = wire.NewSet(
	wire.Struct(new(PostService), "*"),
	wire.Value(10),
	wire.InterfaceValue(new(io.Reader), os.Stdin),
	wire.Bind(new(IPostUsecase), new(*PostUsecase)),
	NewPostUsecase,
	NewPostRepo,
)
