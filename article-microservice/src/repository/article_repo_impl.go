package repository

import (
	"database/sql"
	"errors"
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

func (repo *ArticleRepositoryImpl) Save(article *entity.Article) (*pb.GetArticleResponse, error) {
	// Prepare statement
	const query = `INSERT INTO articles (id_user, title, content) VALUES ($1, $2, $3) RETURNING id, title, content, created_at, updated_at`
	var id int64
	var title string
	var content string
	var created_at time.Time
	var updated_at time.Time

	// Query to db and return id
	err := repo.DB.QueryRow(query, article.AccountId, article.Title, article.Content).Scan(&id, &title, &content, &created_at, &updated_at)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &pb.GetArticleResponse{
		Id:        id,
		Title:     title,
		Content:   content,
		CreatedAt: timestamppb.New(created_at),
		UpdatedAt: timestamppb.New(updated_at),
	}, nil
}

func (repo *ArticleRepositoryImpl) FindById(id int64) (entity.Article, error) {
	const query = `SELECT id_user, title, content, created_at, updated_at, deleted_at FROM articles a WHERE a.id = $1 AND deleted_at IS NULL`
	var title string
	var content string
	var created_at time.Time
	var updated_at time.Time
	var deleted_at sql.NullTime
	var id_user int64

	err := repo.DB.QueryRow(query, id).Scan(
		&id_user, &title, &content, &created_at, &updated_at, &deleted_at)
	if err != nil {
		fmt.Println(err)
		return entity.Article{}, err
	}

	// Check if article has been deleted or not
	if deleted_at.Valid {
		return entity.Article{}, errors.New("article has been deleted")
	}

	return entity.Article{
		Id:        id,
		Title:     title,
		Content:   content,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
		DeletedAt: deleted_at,
		AccountId: id_user,
	}, nil
}

func (repo *ArticleRepositoryImpl) FindAllByUserId(user_id int64) ([]*pb.GetArticleResponse, error) {
	const query = `SELECT ar.id, ar.title, ar.content, ar.created_at, ar.updated_at FROM articles ar INNER JOIN accounts ac ON ar.id_user = ac.id WHERE ac.id = $1 AND ar.deleted_at IS NULL`

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
	const query = `SELECT ar.id, ar.title, ar.content, ar.created_at, ar.updated_at FROM articles ar INNER JOIN accounts ac ON ar.id_user = ac.id WHERE ar.deleted_at IS NULL`

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

func (repo *ArticleRepositoryImpl) Update(article *entity.Article) (*pb.GetArticleResponse, error) {
	query := `UPDATE articles a SET title = $1, content = $2, updated_at = $3 WHERE a.id = $4 AND deleted_at IS NULL RETURNING created_at, updated_at`

	var created_at time.Time
	var updated_at time.Time

	err := repo.DB.QueryRow(query, article.Title, article.Content, time.Now(), article.Id).Scan(&created_at, &updated_at)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &pb.GetArticleResponse{
		Id:        article.Id,
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: timestamppb.New(created_at),
		UpdatedAt: timestamppb.New(updated_at),
	}, nil
}

func (repo *ArticleRepositoryImpl) Delete(id int64) error {
	query := `UPDATE articles a SET deleted_at = $1 WHERE a.id = $2`
	_, err := repo.DB.Exec(query, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}
