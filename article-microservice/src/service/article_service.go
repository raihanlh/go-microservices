package service

import (
	"context"
	"fmt"

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
	// fmt.Println(("HERE"))
	authReq := &pb.GetByTokenRequest{
		Token: req.Token,
	}
	// fmt.Println(authReq.Token)
	var user *pb.GetUserResponse
	user, err := a.AuthService.GetByToken(ctx, authReq)
	fmt.Println(user.Id)

	if err != nil {
		return nil, err
	}
	fmt.Println(user)

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
	// var user *pb.GetUserResponse
	// user, err := a.AuthService.GetByToken(ctx, &pb.GetByTokenRequest{
	// 	Token: req.Token,
	// })

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

// func (a *ArticleServer) GetArticleByUser(ctx context.Context, req *pb.GetArticleRequest)
