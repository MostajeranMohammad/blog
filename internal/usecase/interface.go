package usecase

import (
	"context"

	"github.com/MostajeranMohammad/blog/internal/dto"
	"github.com/MostajeranMohammad/blog/internal/entity"
)

type (
	BlogPost interface {
		CreateNewBlogPost(ctx context.Context, body map[string]interface{}) (entity.BlogPost, error)
		GetBlogPostById(ctx context.Context, id int) (entity.BlogPost, error)
		GetBlogPostByTitle(ctx context.Context, postTitle string) (entity.BlogPost, error)
		GetAllBlogPosts(ctx context.Context, dto dto.FilterBlogPosts, limit, skip int) ([]entity.BlogPost, error)
		Count(ctx context.Context, dto dto.FilterBlogPosts) (int64, error)
		UpdateBlogPost(ctx context.Context, id int, body map[string]interface{}) (entity.BlogPost, error)
		DeletePost(ctx context.Context, id int) (entity.BlogPost, error)
	}
)
