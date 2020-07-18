package protoapi

import (
	"log"
	com "tionyxtrack/masterservice/common"
)

type MasterService struct{}

var ErrorResponse com.IErrorJson

func init() {
	ErrorResponse = com.NewError()
}

func RecoverError() {
	func() {
		if r := recover(); r != nil {
			log.Fatalf("recover from error")
		}
	}()
}

func RecoverErrorMethod(methodName string) {
	func() {
		if r := recover(); r != nil {
			log.Fatalf("recover from error")
		}
	}()
}
