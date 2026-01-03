package applications

import (
	"context"
	"fmt"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/core/types"
	"hinsun-backend/internal/domain/account"
	"hinsun-backend/internal/domain/blog"
	"hinsun-backend/internal/domain/comment"
	"hinsun-backend/internal/domain/usecases"
	"hinsun-backend/internal/domain/values"
)

type BlogAppService interface {
	usecases.ManageBlogUseCase
	usecases.CommentBlogUseCase
}

type blogAppService struct {
	blogService    blog.BlogService
	commentService comment.CommentService
	accountService account.AccountService
}

func NewBlogAppService(
	blogService blog.BlogService,
	commentService comment.CommentService,
	accountService account.AccountService,
) BlogAppService {
	return &blogAppService{
		blogService:    blogService,
		commentService: commentService,
		accountService: accountService,
	}
}

// ================================== ManageBlogUseCase =================================
func (s *blogAppService) FindBlogs(ctx context.Context, query *usecases.FindBlogsQuery) ([]*blog.BlogEntity, error) {
	return s.blogService.FindBlogs(ctx)
}

func (s *blogAppService) CreateBlog(ctx context.Context, params *usecases.CreateBlogParams) (*blog.BlogEntity, error) {
	// 1. Validate and process params
	author, err := s.accountService.FindAccountByID(ctx, params.AuthorID.String())
	if err != nil {
		return nil, err
	}

	if author == nil {
		return nil, failure.NewNotFoundFailure(fmt.Sprintf("author with ID %s not found", params.AuthorID.String()))
	}

	languages, err := values.ConvertStringArrayToMarkdownLanguageCodes(params.Languages)
	if err != nil {
		return nil, err
	}

	// 2. Create blog entity
	blog, err := s.blogService.CreateBlog(
		ctx,
		author.ID,
		languages,
		params.Categories,
		params.Names,
		params.Descriptions,
		params.Markdowns,
		params.IsPublished,
		params.EstimatedReadTimeSeconds,
	)

	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (s *blogAppService) DeleteMultipleBlogs(ctx context.Context, query *usecases.DeleteBlogsQuery) (*types.DeletedResult, error) {
	rowsAffected, err := s.blogService.DeleteMultipleBlogs(ctx, query.IDs)
	if err != nil {
		return nil, err
	}

	return &types.DeletedResult{
		RowsAffected: rowsAffected,
		Payload:      query.IDs,
	}, nil
}

func (s *blogAppService) FindBlog(ctx context.Context, id string) (*blog.BlogEntity, error) {
	return s.blogService.FindBlog(ctx, id)
}

func (s *blogAppService) UpdateBlog(ctx context.Context, id string, params *usecases.UpdateBlogParams) (*blog.BlogEntity, error) {
	author, err := s.accountService.FindAccountByID(ctx, params.AuthorID.String())
	if err != nil {
		return nil, err
	}

	if author == nil {
		return nil, failure.NewNotFoundFailure(fmt.Sprintf("author with ID %s not found", params.AuthorID.String()))
	}

	languages, err := values.ConvertStringArrayToMarkdownLanguageCodes(params.Languages)
	if err != nil {
		return nil, err
	}

	updatedBlog, err := s.blogService.UpdateBlog(
		ctx,
		id,
		author.ID,
		languages,
		params.Categories,
		params.Names,
		params.Descriptions,
		params.Markdowns,
		params.IsPublished,
		params.EstimatedReadTimeSeconds,
	)

	if err != nil {
		return nil, err
	}

	return updatedBlog, nil
}

func (s *blogAppService) DeleteBlog(ctx context.Context, id string) (*types.DeletedResult, error) {
	rowsAffected, err := s.blogService.DeleteBlog(ctx, id)
	if err != nil {
		return nil, err
	}

	return &types.DeletedResult{
		RowsAffected: rowsAffected,
		Payload:      id,
	}, nil
}

// ================================== ManageBlogUseCase =================================

// ================================== CommentBlogUseCase =================================
func (a *blogAppService) AddCommentToBlog(ctx context.Context, blogId string, accountId string, params *usecases.AddCommentToBlogParams) (*comment.CommentEntity, error) {
	return nil, nil
}

func (a *blogAppService) UpdateCommentOnBlog(ctx context.Context, blogId string, accountId string, commentId string, params *usecases.UpdateCommentOnBlogParams) (*comment.CommentEntity, error) {
	return nil, nil
}

func (a *blogAppService) DeleteCommentOnBlog(ctx context.Context, blogId string, accountId string, commentId string) (*types.DeletedResult, error) {
	return nil, nil
}

// ================================== CommentBlogUseCase =================================
