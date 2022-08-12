package Model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResult struct {
	Error string `json:"error"`
}

func NewWrapper[T any, N any](request HandlerType[T, N]) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var reqData T

		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(req.Body).Decode(&reqData)
		if err != nil {
			fmt.Println(err)
			result := ErrorResult{Error: err.Error()}
			jsonData, _ := json.Marshal(result)
			w.Write(jsonData)
			return
		}

		result, err := request.RequestHandler(reqData)

		if err != nil {
			fmt.Println(err)
			result := ErrorResult{Error: err.Error()}
			jsonData, _ := json.Marshal(result)
			w.Write(jsonData)
			return
		}

		jsonData, err := json.Marshal(result)
		w.Write(jsonData)

	}
}

type HandlerType[T any, N any] interface {
	RequestHandler(T) (N, error)
}

// Deprecated: Should not be used anymore
type HandlerResult[N any] struct {
	Result N
	Error  error
}
