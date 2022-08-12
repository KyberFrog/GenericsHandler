package TestHandler

import (
	"testing"
)

func TestHandlerCase1(t *testing.T) {
	h := TestHandler[ReqData, ResData]{}
	req := ReqData{Username: "admin"}
	res, err := h.RequestHandler(req)

	if err != nil {
		t.Error("Error should be nil!")
	}

	if res.Result != true {
		t.Error("Result should be true!")
	}
}

func TestHandlerCase2(t *testing.T) {
	h := TestHandler[ReqData, ResData]{}
	req := ReqData{Username: "test123"}
	_, err := h.RequestHandler(req)

	if err == nil {
		t.Error("Error should not be nil!")
	}

}
