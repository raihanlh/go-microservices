package service

import (
	"context"

	pb "github.com/raihanlh/go-article-microservice/proto"
	"github.com/raihanlh/go-article-microservice/src/entity"
	"github.com/raihanlh/go-article-microservice/src/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	return &pb.CreateArticleResponse{
		Article: article,
		Message: "Article created succesfully",
		Status:  "201",
	}, nil
}

func (a *ArticleServer) GetArticleById(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {
	article, err := a.ArticleRepository.FindById(req.Id)
	if (err != nil) || (article.DeletedAt.Valid) {
		return &pb.GetArticleResponse{}, status.Error(codes.NotFound, "article not found")
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
	var user *pb.GetUserResponse
	user, err := a.AuthService.GetByToken(ctx, &pb.GetByTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return nil, err
	}

	article, err := a.ArticleRepository.FindById(req.Id)
	if err != nil {
		return nil, err
	}
	if article.AccountId != user.Id {
		return nil, status.Error(codes.PermissionDenied, "unauthorized")
	}

	res, err := a.ArticleRepository.Update(&entity.Article{
		Id:        req.Id,
		Title:     req.Title,
		Content:   req.Content,
		AccountId: user.Id,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *ArticleServer) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error) {
	var user *pb.GetUserResponse
	user, err := a.AuthService.GetByToken(ctx, &pb.GetByTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return nil, err
	}

	article, err := a.ArticleRepository.FindById(req.Id)
	if err != nil {
		return nil, err
	}
	if article.AccountId != user.Id {
		return nil, status.Error(codes.PermissionDenied, "unauthorized")
	}

	err = a.ArticleRepository.Delete(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteArticleResponse{
		Status:  "201",
		Message: "Success",
	}, nil
}
