package repo

import (
	"context"

	"github.com/MostajeranMohammad/blog/internal/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BlogPostRepo struct {
	db *gorm.DB
}

func NewBlogPostRepo(db *gorm.DB) BlogPost {
	return &BlogPostRepo{
		db,
	}
}

func (br *BlogPostRepo) Create(ctx context.Context, p map[string]interface{}) (entity.BlogPost, error) {
	createdPost := entity.BlogPost{}
	result := br.db.Model(&createdPost).WithContext(ctx).Clauses(clause.Returning{}).Create(p)
	return createdPost, result.Error
}

func (br *BlogPostRepo) GetById(ctx context.Context, id int) (entity.BlogPost, error) {
	blogPost := entity.BlogPost{}
	result := br.db.WithContext(ctx).First(&blogPost)
	return blogPost, result.Error
}

func (br *BlogPostRepo) GetByTitle(ctx context.Context, title string) (entity.BlogPost, error) {
	blogPost := entity.BlogPost{Title: title}
	result := br.db.WithContext(ctx).First(&blogPost)
	return blogPost, result.Error
}

func (br *BlogPostRepo) GetAll(ctx context.Context, filter clause.AndConditions, skip int, limit int) ([]entity.BlogPost, error) {
	blogPosts := []entity.BlogPost{}
	if limit == 0 {
		limit = 10
	}

	query := br.db.WithContext(ctx).Limit(limit).Offset(skip)
	if len(filter.Exprs) > 0 {
		query.Where(filter)
	}

	result := query.Find(&blogPosts)
	return blogPosts, result.Error
}

func (br *BlogPostRepo) Update(ctx context.Context, id int, m map[string]interface{}) (entity.BlogPost, error) {
	updatedBlogPost := entity.BlogPost{}
	result := br.db.WithContext(ctx).Clauses(clause.Returning{}).Model(&updatedBlogPost).Where("id = ?", id).Updates(m)
	return updatedBlogPost, result.Error
}

func (br *BlogPostRepo) Delete(ctx context.Context, id int) (entity.BlogPost, error) {
	deletedPost := entity.BlogPost{}
	result := br.db.WithContext(ctx).Clauses(clause.Returning{}).Where("id = ?", id).Delete(&deletedPost)
	return deletedPost, result.Error
}
