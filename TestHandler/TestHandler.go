package TestHandler

import (
	"errors"
	"genericsTest/Model"
)

type TestHandler[T ReqData, N ResData] struct{}

func (receiver TestHandler[T, N]) RequestHandler(data ReqData) (N, error) {
	if data.Username != "admin" {
		return N{}, errors.New("Unknown user")
	}

	return N{Result: true}, nil
}

type ReqData struct {
	Username string `json:"username"`
}

type ResData struct {
	Result bool `json:"result"`
}

var handler = TestHandler[ReqData, ResData]{}
var SharedWrapper = Model.NewWrapper[ReqData, ResData](handler)
