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

// PostUsecaseOption PostUsecaseOption
type PostUsecaseOption struct {
	a    A
	b    B
	repo IPostRepo
}

// NewPostUsecase NewPostUsecase
func NewPostUsecase(opt *PostUsecaseOption) (*PostUsecase, func(), error) {
	return &PostUsecase{repo: opt.repo}, nil, nil
}

// service service

type A int
type B int

// PostService PostService
type PostService struct {
	usecase IPostUsecase
	a       A
	b       B
	r       io.Reader
}

// NewPostService NewPostService
func NewPostService(u IPostUsecase) (*PostService, error) {
	return &PostService{usecase: u}, nil
}

// PostServiceSet PostServiceSet
var PostServiceSet = wire.NewSet(
	wire.Struct(new(PostService), "*"),
	wire.Value(A(10)),
	wire.Value(B(10)),
	wire.InterfaceValue(new(io.Reader), os.Stdin),

	// for usecase
	wire.Bind(new(IPostUsecase), new(*PostUsecase)),
	wire.Struct(new(PostUsecaseOption), "*"),
	NewPostUsecase,

	NewPostRepo,
)
