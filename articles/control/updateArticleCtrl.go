package control

import (
	"errors"

	"example.com/rest/articles/models"
)

type UpdateArticleRequestBody struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type UpdateArticleApiResponseItem struct {
	models.Article
}

type UpdateArticleApiResponse struct {
	Item            *UpdateArticleApiResponseItem
	Error           error
	ErrorStatusCode int
	ErrorAppCode    int
}

func (ctrl *UpdateArticleRequestBody) verify() error {
	logger.Info("request validation/verification start")
	defer logger.Info("request validation/verification end")
	if ctrl.Id == "" {
		return errors.New("Id is required")
	}
	if ctrl.Description == "" {
		return errors.New("Description is required")
	}
	//check id is available
	if _, found := models.ArticleList[ctrl.Id]; !found {
		return errors.New("Id is not available")
	}

	return nil
}

func (ctrl *UpdateArticleRequestBody) Run() *UpdateArticleApiResponse {
	if err := ctrl.verify(); err != nil {
		return &UpdateArticleApiResponse{
			Error: err,
		}
	}
	return ctrl.update()
}

func (ctrl *UpdateArticleRequestBody) update() *UpdateArticleApiResponse {
	article := models.Article{
		Id:          ctrl.Id,
		Title:       models.ArticleList[ctrl.Id].Title,
		Description: ctrl.Description,
	}

	logger.Info("update article...")
	models.ArticleList[article.Id] = article

	return &UpdateArticleApiResponse{
		Item: &UpdateArticleApiResponseItem{article},
	}
}
