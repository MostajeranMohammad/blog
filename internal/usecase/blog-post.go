package usecase

import (
	"context"

	"github.com/MostajeranMohammad/blog/internal/dto"
	"github.com/MostajeranMohammad/blog/internal/entity"
	"github.com/MostajeranMohammad/blog/internal/repo"
)

type BlogPostUsecase struct {
	repo repo.BlogPost
}

func NewBlogPostUsecase(repo repo.BlogPost) BlogPost {
	return &BlogPostUsecase{repo}
}

func (bu *BlogPostUsecase) CreateNewBlogPost(ctx context.Context, body map[string]interface{}) (entity.BlogPost, error) {
	return bu.repo.Create(ctx, body)
}

func (bu *BlogPostUsecase) GetBlogPostById(ctx context.Context, id int) (entity.BlogPost, error) {
	return bu.repo.GetById(ctx, id)
}

func (bu *BlogPostUsecase) GetBlogPostByTitle(ctx context.Context, postTitle string) (entity.BlogPost, error) {
	return bu.repo.GetByTitle(ctx, postTitle)
}

func (bu *BlogPostUsecase) GetAllBlogPosts(ctx context.Context, dto dto.FilterBlogPosts, limit, skip int) ([]entity.BlogPost, error) {
	query, err := dto.ToQueryModel()
	if err != nil {
		return nil, err
	}
	return bu.repo.GetAll(ctx, query, skip, limit)
}

func (bu *BlogPostUsecase) UpdateBlogPost(ctx context.Context, id int, body map[string]interface{}) (entity.BlogPost, error) {
	return bu.repo.Update(ctx, id, body)
}

func (bu *BlogPostUsecase) DeletePost(ctx context.Context, id int) (entity.BlogPost, error) {
	return bu.repo.Delete(ctx, id)
}
