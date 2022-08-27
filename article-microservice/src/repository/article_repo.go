package repository

import (
	pb "github.com/raihanlh/go-article-microservice/proto"
	"github.com/raihanlh/go-article-microservice/src/entity"
)

type ArticleRepository interface {
	Save(article *entity.Article) (*pb.GetArticleResponse, error)
	Update(article *entity.Article) (*pb.GetArticleResponse, error)
	Delete(id int64) error
	FindById(id int64) (entity.Article, error)
	FindAllByUserId(user_id int64) ([]*pb.GetArticleResponse, error)
	FindAll() ([]*pb.GetArticleResponse, error)
}
