package articles

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/rest/articles/control"
	"go.uber.org/zap"
)

func deleteArticlesHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("delete article request start")
	defer logger.Info("delete article request stop")

	var body control.DeleteArticleRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Error("error while decoding article request body into struct,", zap.Error(err))
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		data := fmt.Sprintf(`{"status": "error", "message": "bad request invalid json data/ invalid format"}`)
		w.Write([]byte(data))
		return
	}

	data := body.Run()

	if data.Error != nil {
		data.ErrorStatusCode = 400
		w.WriteHeader(data.ErrorStatusCode)
		data := fmt.Sprintf(`{"status": "error", "message": "%s"}`, data.Error.Error())
		w.Write([]byte(data))
		return
	}

	raw, _ := json.Marshal(data.Items)

	w.Write(raw)

}
