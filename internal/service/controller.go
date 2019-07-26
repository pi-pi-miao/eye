package service

import (
	"io"
	"net/http"
	"eye/pkg/logger"
	"github.com/julienschmidt/httprouter"
)

func GetHello(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	logger.Debugf("method is %v request remote %v",r.Method,r.RemoteAddr)
	io.WriteString(w,"hello world")
	return
}
