package usecase

import (
	"github.com/MostajeranMohammad/blog/internal/dto"
	"github.com/MostajeranMohammad/blog/internal/entity"
)

type (
	BlogPost interface {
		CreateNewBlogPost(body dto.CreateNewBlogPostDto) (entity.BlogPost, error)
		GetBlogPostById(id int) (entity.BlogPost, error)
		GetBlogPostByTitle(postTitle string) (entity.BlogPost, error)
		GetAllBlogPosts(dto dto.FilterBlogPosts, limit, skip int) ([]entity.BlogPost, error)
		UpdateBlogPost(id int, body dto.UpdateBlogPostDto) (entity.BlogPost, error)
		DeletePost(id int) (entity.BlogPost, error)
	}
)
