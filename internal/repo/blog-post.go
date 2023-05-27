package repo

import (
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

func (br *BlogPostRepo) Create(p entity.BlogPost) (entity.BlogPost, error) {
	result := br.db.Clauses(clause.Returning{}).Create(&p)
	return p, result.Error
}

func (br *BlogPostRepo) GetById(id int) (entity.BlogPost, error) {
	blogPost := entity.BlogPost{}
	result := br.db.First(&blogPost)
	return blogPost, result.Error
}

func (br *BlogPostRepo) GetByTitle(title string) (entity.BlogPost, error) {
	blogPost := entity.BlogPost{Title: title}
	result := br.db.First(&blogPost)
	return blogPost, result.Error
}

func (br *BlogPostRepo) GetAll(filter clause.AndConditions, skip int, limit int) ([]entity.BlogPost, error) {
	blogPosts := []entity.BlogPost{}
	if limit == 0 {
		limit = 10
	}

	query := br.db.Limit(limit).Offset(skip)
	if len(filter.Exprs) > 0 {
		query.Where(filter)
	}

	result := query.Find(&blogPosts)
	return blogPosts, result.Error
}

func (br *BlogPostRepo) Update(id int, m entity.BlogPost) (entity.BlogPost, error) {
	updatedBlogPost := entity.BlogPost{}
	result := br.db.Clauses(clause.Returning{}).Model(&updatedBlogPost).Where("id = ?", id).Updates(m)
	return updatedBlogPost, result.Error
}

func (br *BlogPostRepo) Delete(id int) (entity.BlogPost, error) {
	deletedPost := entity.BlogPost{}
	result := br.db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&deletedPost)
	return deletedPost, result.Error
}
