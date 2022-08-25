package repository

import (
	"github.com/raihanlh/go-article-microservice/src/entity"
)

type ArticleRepository interface {
	Save(Article *entity.Article) (entity.Article, error)
	Update(Article *entity.Article) (entity.Article, error)
	FindById(id int64) (entity.Article, error)
	FindAllByUserId(user_id int64) ([]entity.Article, error)
}
