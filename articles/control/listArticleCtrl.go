package control

import (
	"example.com/rest/articles/models"
)

type ArticleListApiResponse struct {
	Items           []models.Article `json:"items"`
	Error           error            `json:"error"`
	ErrorStatusCode int              `json:"-"`
	ErrorAppCode    int              `json:"code,omitempty"`
}

func Run() *ArticleListApiResponse {
	return listSamiti()
}

func listSamiti() *ArticleListApiResponse {

	logger.Info("Reading article from slice...")

	return &ArticleListApiResponse{
		Items: ArticleList,
	}
}
