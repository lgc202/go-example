package example

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
type postUsecase struct {
	repo IPostRepo
}

// NewPostUsecase NewPostUsecase
func NewPostUsecase(repo IPostRepo) (IPostUsecase, func(), error) {
	return postUsecase{repo: repo}, nil, nil
}

// service service

// PostService PostService
type PostService struct {
	usecase IPostUsecase
}

// NewPostService NewPostService
func NewPostService(u IPostUsecase) (*PostService, error) {
	return &PostService{usecase: u}, nil
}
