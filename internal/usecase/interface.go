package usecase

import (
	"context"

	"github.com/MostajeranMohammad/blog/internal/dto"
	"github.com/MostajeranMohammad/blog/internal/entity"
)

type (
	BlogPost interface {
		CreateNewBlogPost(ctx context.Context, body dto.CreateNewBlogPostDto) (entity.BlogPost, error)
		GetBlogPostById(ctx context.Context, id int) (entity.BlogPost, error)
		GetBlogPostByTitle(ctx context.Context, postTitle string) (entity.BlogPost, error)
		GetAllBlogPosts(ctx context.Context, dto dto.FilterBlogPosts, limit, skip int) ([]entity.BlogPost, error)
		UpdateBlogPost(ctx context.Context, id int, body dto.UpdateBlogPostDto) (entity.BlogPost, error)
		DeletePost(ctx context.Context, id int) (entity.BlogPost, error)
	}
)
