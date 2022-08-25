package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/raihanlh/go-article-microservice/src/entity"
)

type ArticleRepositoryImpl struct {
	DB *sql.DB
}

func NewArticleRepository(db *sql.DB) ArticleRepository {
	return &ArticleRepositoryImpl{
		DB: db,
	}
}

func (repo *ArticleRepositoryImpl) Save(article *entity.Article) (entity.Article, error) {
	// Prepare statement
	const query = `INSERT INTO articles (id_user, title, content) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	var id int64
	var created_at time.Time
	var updated_at time.Time

	// Query to db and return id
	err := repo.DB.QueryRow(query, article.AccountId, article.Title, article.Content).Scan(&id, &created_at, &updated_at)
	if err != nil {
		fmt.Println(err.Error())
		return entity.Article{}, err
	}

	return entity.Article{
		Id:        id,
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}, nil
}

func (repo *ArticleRepositoryImpl) Update(Article *entity.Article) (entity.Article, error) {
	return entity.Article{}, nil
}
func (repo *ArticleRepositoryImpl) GetById(id int64) (entity.Article, error) {
	return entity.Article{}, nil
}
func (repo *ArticleRepositoryImpl) GetAllByUserId(user_id int64) ([]entity.Article, error) {
	return make([]entity.Article, 1), nil
}
