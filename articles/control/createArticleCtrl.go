package control

import (
	"errors"

	"example.com/rest/articles/models"

	"github.com/google/uuid"
)

type CreateArticleRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateArticleApiResponseItem struct {
	models.Article
}

type CreateArticleApiResponse struct {
	Item            *CreateArticleApiResponseItem
	Error           error
	ErrorStatusCode int
	ErrorAppCode    int
}

func (ctrl *CreateArticleRequestBody) verify() error {
	logger.Info("request validation/verification start")
	defer logger.Info("request validation/verification end")
	if ctrl.Title == "" {
		return errors.New("Title is required")
	}
	if ctrl.Description == "" {
		return errors.New("Description is required")
	}
	return nil
}

func (ctrl *CreateArticleRequestBody) Run() *CreateArticleApiResponse {
	if err := ctrl.verify(); err != nil {
		return &CreateArticleApiResponse{
			Error: err,
		}
	}
	return ctrl.create()
}

func (ctrl *CreateArticleRequestBody) create() *CreateArticleApiResponse {
	article := models.Article{
		Id:          uuid.NewString(),
		Title:       ctrl.Title,
		Description: ctrl.Description,
	}

	logger.Info("saving article...")
	models.ArticleList[article.Id] = article
	logger.Info("saved article...")
	return &CreateArticleApiResponse{
		Item: &CreateArticleApiResponseItem{article},
	}
}
