package service

import (
	"context"

	pb "github.com/raihanlh/go-article-microservice/proto"
	"github.com/raihanlh/go-article-microservice/src/entity"
	"github.com/raihanlh/go-article-microservice/src/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ArticleServer struct {
	pb.UnimplementedArticleServiceServer
	ArticleRepository repository.ArticleRepository
	AuthService       pb.AuthServiceClient
}

func (a *ArticleServer) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	authReq := &pb.GetByTokenRequest{
		Token: req.Token,
	}
	var user *pb.GetUserResponse
	user, err := a.AuthService.GetByToken(ctx, authReq)

	if err != nil {
		return nil, err
	}

	article, err := a.ArticleRepository.Save(&entity.Article{
		AccountId: user.Id,
		Title:     req.Title,
		Content:   req.Content,
	})

	if err != nil {
		return nil, err
	}

	res := pb.GetArticleResponse{
		Id:        article.Id,
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: timestamppb.New(article.CreatedAt),
		UpdatedAt: timestamppb.New(article.UpdatedAt),
	}

	return &pb.CreateArticleResponse{
		Article: &res,
		Message: "Article created succesfully",
		Status:  "201",
	}, nil
}

func (a *ArticleServer) GetArticleById(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {
	article, err := a.ArticleRepository.FindById(req.Id)
	if err != nil {
		return &pb.GetArticleResponse{}, err
	}

	return &pb.GetArticleResponse{
		Id:        article.Id,
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: timestamppb.New(article.CreatedAt),
		UpdatedAt: timestamppb.New(article.UpdatedAt),
	}, nil
}

func (a *ArticleServer) GetArticleByUser(ctx context.Context, req *pb.GetAllArticleByUserRequest) (*pb.GetAllArticleResponse, error) {
	var user *pb.GetUserResponse
	user, err := a.AuthService.GetByToken(ctx, &pb.GetByTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return nil, err
	}

	articles, err := a.ArticleRepository.FindAllByUserId(user.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetAllArticleResponse{
		Articles: articles,
	}, nil

}

func (a *ArticleServer) GetAllArticle(ctx context.Context, req *pb.GetAllArticleRequest) (*pb.GetAllArticleResponse, error) {
	articles, err := a.ArticleRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return &pb.GetAllArticleResponse{
		Articles: articles,
	}, nil

}

func (a *ArticleServer) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.GetArticleResponse, error) {
	return &pb.GetArticleResponse{}, nil
}
