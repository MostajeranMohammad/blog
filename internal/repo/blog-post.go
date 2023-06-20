package repo

import (
	"context"
	"time"

	"github.com/MostajeranMohammad/blog/internal/entity"
	"github.com/MostajeranMohammad/blog/pkg/utils"
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

func (br *BlogPostRepo) Create(ctx context.Context, d map[string]interface{}) (entity.BlogPost, error) {
	post := utils.ChangeMapFieldsToSnakeCase(d)
	post["created_at"] = time.Now()
	result := br.db.Model(entity.BlogPost{}).WithContext(ctx).Clauses(clause.Returning{}).Create(&post)
	return br.mapToStruct(post), result.Error
}

func (br *BlogPostRepo) GetById(ctx context.Context, id int) (entity.BlogPost, error) {
	blogPost := entity.BlogPost{}
	result := br.db.WithContext(ctx).Where("id = ?", id).First(&blogPost)
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

func (br *BlogPostRepo) Update(ctx context.Context, id int, d map[string]interface{}) (entity.BlogPost, error) {
	post := utils.ChangeMapFieldsToSnakeCase(d)
	tempPost := entity.BlogPost{}
	result := br.db.WithContext(ctx).Clauses(clause.Returning{}).Model(&tempPost).Where("id = ?", id).Updates(post)
	return tempPost, result.Error
}

func (br *BlogPostRepo) Delete(ctx context.Context, id int) (entity.BlogPost, error) {
	deletedPost := entity.BlogPost{}
	result := br.db.WithContext(ctx).Clauses(clause.Returning{}).Where("id = ?", id).Delete(&deletedPost)
	return deletedPost, result.Error
}

func (br *BlogPostRepo) mapToStruct(m map[string]interface{}) entity.BlogPost {
	return entity.BlogPost{
		Title:    m["title"].(string),
		AuthorId: m["author_id"].(uint),
		Content:  m["content"].(string),
		Model: gorm.Model{
			ID: m["id"].(uint),
		},
	}
}
