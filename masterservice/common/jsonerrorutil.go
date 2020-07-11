package common

import pro "tionyxtrack/masterservice/proto"
import "encoding/json"

type IErrorJson interface {
	GetCreateErrorJson(err error) *pro.ResponseError
}
type ErrorJson struct{}

func NewError() IErrorJson {
	return &ErrorJson{}
}
func (e *ErrorJson) GetCreateErrorJson(err error) *pro.ResponseError {
	errorResponse := &pro.ResponseError{}
	if err != nil {
		jsonError, _ := json.Marshal(err)
		errorResponse.ErrorMessage = jsonError
		errorResponse.IsError = true
	} else {
		errorResponse.ErrorMessage = make([]byte, 0)
		errorResponse.IsError = true
	}
	return errorResponse
}
