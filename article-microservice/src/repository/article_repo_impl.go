package repository

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/raihanlh/go-article-microservice/proto"
	"github.com/raihanlh/go-article-microservice/src/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	// result, err := repo.DB.Exec(query, article.AccountId, article.Title, article.Content).Scan(&id, &created_at, &updated_at)
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
func (repo *ArticleRepositoryImpl) FindById(id int64) (entity.Article, error) {
	const query = `SELECT title, content, created_at, updated_at FROM articles a WHERE a.id = $1`
	var title string
	var content string
	var created_at time.Time
	var updated_at time.Time

	err := repo.DB.QueryRow(query, id).Scan(
		&title, &content, &created_at, &updated_at)
	if err != nil {
		fmt.Println(err)
		return entity.Article{}, err
	}

	return entity.Article{
		Id:        id,
		Title:     title,
		Content:   content,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}, nil
}

func (repo *ArticleRepositoryImpl) FindAllByUserId(user_id int64) ([]*pb.GetArticleResponse, error) {
	const query = `select ar.id, ar.title, ar.content, ar.created_at, ar.updated_at FROM articles ar INNER JOIN accounts ac ON ar.id_user = ac.id WHERE ac.id = $1`

	articles := make([]*pb.GetArticleResponse, 0)
	rows, err := repo.DB.Query(query, user_id)
	if err != nil {
		return make([]*pb.GetArticleResponse, 0), err
	}

	for rows.Next() {
		var id int64
		var title string
		var content string
		var created_at time.Time
		var updated_at time.Time

		err = rows.Scan(&id, &title, &content, &created_at, &updated_at)

		if err != nil {
			return make([]*pb.GetArticleResponse, 0), err
		}

		articles = append(articles, &pb.GetArticleResponse{
			Id:        id,
			Title:     title,
			Content:   content,
			CreatedAt: timestamppb.New(created_at),
			UpdatedAt: timestamppb.New(updated_at),
		})
	}

	return articles, nil
}

func (repo *ArticleRepositoryImpl) FindAll() ([]*pb.GetArticleResponse, error) {
	const query = `select ar.id, ar.title, ar.content, ar.created_at, ar.updated_at FROM articles ar INNER JOIN accounts ac ON ar.id_user = ac.id`

	articles := make([]*pb.GetArticleResponse, 0)
	rows, err := repo.DB.Query(query)
	if err != nil {
		return make([]*pb.GetArticleResponse, 0), err
	}

	for rows.Next() {
		var id int64
		var title string
		var content string
		var created_at time.Time
		var updated_at time.Time

		err = rows.Scan(&id, &title, &content, &created_at, &updated_at)

		if err != nil {
			return make([]*pb.GetArticleResponse, 0), err
		}

		articles = append(articles, &pb.GetArticleResponse{
			Id:        id,
			Title:     title,
			Content:   content,
			CreatedAt: timestamppb.New(created_at),
			UpdatedAt: timestamppb.New(updated_at),
		})
	}

	return articles, nil
}
