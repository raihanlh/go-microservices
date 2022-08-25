package repository

import (
	"github.com/raihanlh/go-article-microservice/src/entity"
)

type ArticleRepository interface {
	Save(Article *entity.Article) (entity.Article, error)
	Update(Article *entity.Article) (entity.Article, error)
	GetById(id int64) (entity.Article, error)
	GetAllByUserId(user_id int64) ([]entity.Article, error)
}
