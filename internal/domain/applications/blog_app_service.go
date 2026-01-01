package applications

import (
	"context"
	"hinsun-backend/internal/core/types"
	"hinsun-backend/internal/domain/blog"
	"hinsun-backend/internal/domain/usecases"
)

type BlogAppSevice interface {
	usecases.ManageBlogUseCase
}

type blogAppService struct {
	blogService blog.BlogService
}

func NewBlogAppService(
	blogService blog.BlogService,
) BlogAppSevice {
	return &blogAppService{
		blogService: blogService,
	}
}

func (s *blogAppService) FindBlogs(ctx context.Context, query *usecases.FindBlogsQuery) ([]*blog.BlogEntity, error) {
	return nil, nil
}

func (s *blogAppService) CreateBlog(ctx context.Context, params *usecases.CreateBlogParams) (*blog.BlogEntity, error) {
	return nil, nil
}

func (s *blogAppService) DeleteMultipleBlogs(ctx context.Context, query *usecases.DeleteBlogsQuery) (*types.DeletedResult, error) {
	return nil, nil
}

func (s *blogAppService) FindBlog(ctx context.Context, id string) (*blog.BlogEntity, error) {
	return nil, nil
}

func (s *blogAppService) UpdateBlog(ctx context.Context, id string, params *usecases.UpdateBlogParams) (*blog.BlogEntity, error) {
	return nil, nil
}

func (s *blogAppService) DeleteBlog(ctx context.Context, id string) (*types.DeletedResult, error) {
	return nil, nil
}
