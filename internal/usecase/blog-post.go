package usecase

import (
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

func (bu *BlogPostUsecase) CreateNewBlogPost(body dto.CreateNewBlogPostDto) (entity.BlogPost, error) {
	return bu.repo.Create(body.ToEntityModel())
}

func (bu *BlogPostUsecase) GetBlogPostById(id int) (entity.BlogPost, error) {
	return bu.repo.GetById(id)
}

func (bu *BlogPostUsecase) GetBlogPostByTitle(postTitle string) (entity.BlogPost, error) {
	return bu.repo.GetByTitle(postTitle)
}

func (bu *BlogPostUsecase) GetAllBlogPosts(dto dto.FilterBlogPosts, limit, skip int) ([]entity.BlogPost, error) {
	query, err := dto.ToQueryModel()
	if err != nil {
		return nil, err
	}
	return bu.repo.GetAll(query, skip, limit)
}

func (bu *BlogPostUsecase) UpdateBlogPost(id int, body dto.UpdateBlogPostDto) (entity.BlogPost, error) {
	return bu.repo.Update(id, body.ToEntityModel())
}

func (bu *BlogPostUsecase) DeletePost(id int) (entity.BlogPost, error) {
	return bu.repo.Delete(id)
}
