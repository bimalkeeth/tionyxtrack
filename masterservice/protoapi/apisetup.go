package protoapi

import com "tionyxtrack/masterservice/common"

type MasterService struct{}

var ErrorResponse com.IErrorJson

func init() {
	ErrorResponse = com.NewError()
}
