package main

import (
	"github.com/gorilla/mux"
	"log"
	"sync"
)

type ObjectAPIHandlers struct {
	ObjectAPI ObjectLayer
}

type ObjectLayer interface {
	// Bucket operations.
	MakeBucket() error
}

func (y *YigStorage) MakeBucket() {
	//TODO implement me

}

type YigStorage struct {
	Logger    log.Logger
	Stopping  bool
	WaitGroup *sync.WaitGroup
}

func RegisterAPIRouter(router *mux.Router, api ObjectAPIHandlers) {
	// API Router
	// 设置路径前缀和子路由
	//apiRouter := router.PathPrefix("/").Subrouter()
	//apiRouter.Methods("get").HandlerFunc(api.ObjectAPI.MakeBucket())
}

type ServerConfig struct {
	Address      string
	KeyFilePath  string     // path for SSL key file
	CertFilePath string     // path for SSL certificate file
	Logger       log.Logger // global logger
	ObjectLayer  *YigStorage
}

// configureServer handler returns final handler for the http server.
func configureServerHandler(c *ServerConfig) {

	return
}
