package control

import (
	"errors"

	"example.com/rest/articles/models"
)

type DeleteArticleRequestBody struct {
	Id string `json:"id"`
}

type DeleteArticleApiResponse struct {
	Items           map[string]models.Article
	Error           error
	ErrorStatusCode int
	ErrorAppCode    int
}

func (ctrl *DeleteArticleRequestBody) verify() error {
	logger.Info("request validation/verification start")
	defer logger.Info("request validation/verification end")
	if ctrl.Id == "" {
		return errors.New("Id is required")
	}
	if _, found := models.ArticleList[ctrl.Id]; !found {
		return errors.New("Id is not available")
	}

	return nil
}

func (ctrl *DeleteArticleRequestBody) Run() *DeleteArticleApiResponse {
	if err := ctrl.verify(); err != nil {
		return &DeleteArticleApiResponse{
			Error: err,
		}
	}
	return ctrl.delete()
}

func (ctrl *DeleteArticleRequestBody) delete() *DeleteArticleApiResponse {
	delete(models.ArticleList, ctrl.Id)

	return &DeleteArticleApiResponse{
		Items: models.ArticleList,
	}
}
