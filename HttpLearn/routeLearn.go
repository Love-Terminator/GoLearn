package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Grantee struct {
	XMLName      xml.Name `xml:"Grantee"`
	XmlnsXsi     string   `xml:"xmlns:xsi,attr"`
	XsiType      string   `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	URI          string   `xml:"URI,omitempty"`
	ID           string   `xml:"ID,omitempty"`
	DisplayName  string   `xml:"DisplayName,omitempty"`
	EmailAddress string   `xml:"EmailAddress,omitempty"`
}

type Grant struct {
	XMLName    xml.Name `xml:"Grant"`
	Grantee    Grantee  `xml:"Grantee"`
	Permission string   `xml:"Permission"`
}

type Acl struct {
	CannedAcl string  `json:"CannedAcl,omitempty"`
	GrantList []Grant `json:"GrantList,omitempty"`
}

func SetupConfig() {
	fmt.Println("Set Config")
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "我是Admin")
	// acl := new(Acl)
	// fmt.Println(*acl)
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(r.Context())
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(context.Background())
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(r.Header)
	fmt.Fprintln(w, "我是Api")

}

func RegisterAdminHandler(route *mux.Router) {
	adminRoute := route.PathPrefix("/admin").Subrouter()

	bucket := adminRoute.Host("s3.test.com").PathPrefix("/get/info").Subrouter()
	bucket.Methods("get").HandlerFunc(AdminHandler)
}

func RegisterApiHandler(route *mux.Router) {
	adminRoute := route.PathPrefix("/api").Subrouter()

	bucket := adminRoute.Host("s3.test.com").PathPrefix("/put/bucket/{bucket}").Subrouter()
	bucket.Methods("get").HandlerFunc(ApiHandler)
}

func startApiRoute(wg *sync.WaitGroup) {
	apiRoute := mux.NewRouter()
	RegisterApiHandler(apiRoute)
	server := &http.Server{
		Addr:    ":8000",
		Handler: apiRoute,
	}
	server.ListenAndServe()
	/*
		wg.Add(1)
		go func() {
			defer wg.Done()
			var err error
			// Configure TLS if certs are available.
			err = server.ListenAndServe()
			if err != http.ErrServerClosed {
				fmt.Println("API server error.")
			}
		}()

	*/
}

func startAdminRoute(wg *sync.WaitGroup) {

	adminRoute := mux.NewRouter()
	RegisterAdminHandler(adminRoute)
	server := &http.Server{
		Addr:    ":8080",
		Handler: adminRoute,
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
	/*
		wg.Add(1)
		go func() {
			defer wg.Done()
			var err error
			// Configure TLS if certs are available.
			err = server.ListenAndServe()
			if err != http.ErrServerClosed {
				fmt.Println("Admin server error.")
			}
		}()

	*/
}

type handlerFunc func(http.Handler) http.Handler

var handlerFns = []handlerFunc{
	//	SetJwtMiddlewareHandler,
}

func RegisterHandlers(router *mux.Router, handlerFns ...handlerFunc) http.Handler {
	var f http.Handler
	f = router
	for _, hFn := range handlerFns {
		f = hFn(f)
	}
	return f
}

func main() {
	wg := &sync.WaitGroup{}
	startAdminRoute(wg)
	startApiRoute(wg)

	signal.Ignore()
	signalQueue := make(chan os.Signal)
	signal.Notify(signalQueue,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGHUP)

	for {
		s := <-signalQueue
		switch s {
		case syscall.SIGHUP:
			// reload config file
			SetupConfig()
		default:
			// stop YIG server, order matters
			wg.Wait()
			return
		}
	}
}
