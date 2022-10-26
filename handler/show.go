package handler

import (
	"fmt"
	"net/http"

	"github.com/yezige/go-short/logx"
)

type any interface{}

func IsSuccess(w http.ResponseWriter, data string) {
	logx.LogAccess.Infoln("Success: " + data[:3])
	fmt.Fprint(w, data)
}

func IsError(w http.ResponseWriter, err any) {
	errstr := ""
	switch err := err.(type) {
	case error:
		errstr = err.Error()
	case string:
		errstr = err
	}
	logx.LogAccess.Infoln("Error: " + errstr)
	fmt.Fprint(w, errstr)
}
