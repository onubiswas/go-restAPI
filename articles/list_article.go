package articles

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/rest/articles/control"
)

func listArticlesHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("list of article request start")
	defer logger.Info("list of  article request stop")

	data := control.Run()
	w.Header().Add("content-type", "application/json")

	if data.Error != nil {
		data.ErrorStatusCode = 400
		w.WriteHeader(data.ErrorStatusCode)
		data := fmt.Sprintf(`{"status": "error", "message": "%s"}`, data.Error.Error())
		w.Write([]byte(data))
		return
	}

	raw, _ := json.Marshal(data.Items)
	msg := "List of articles: "
	w.Write([]byte(msg))
	w.Write(raw)
}
