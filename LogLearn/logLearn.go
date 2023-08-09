package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"sync"
)

// muxDemo5。中间件进阶测试。中间件CommonHeaderHandler和LogHeaderHandler会被串起来。在执行Helloworld之前执行
type HandlerFunc func(http.Handler) http.Handler

type commonHeaderHandler struct {
	handler http.Handler
}

func (h commonHeaderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("我是commonHeaderHandler"))
	w.Header().Set("commonHeader", "common")
	fmt.Println(w.Header())
	if err != nil {
		log.Error(err)
	}
	h.handler.ServeHTTP(w, r)
	log.Info("我是commonHeaderHandler")
}

func SetCommonHeaderHandler(h http.Handler) http.Handler {
	return commonHeaderHandler{h}
}

type logHeaderHandler struct {
	handler http.Handler
}

func (h logHeaderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("我是logHeaderHandler"))
	w.Header().Set("logHeader", "log")
	fmt.Println(w.Header())
	if err != nil {
		log.Error(err)
	}
	h.handler.ServeHTTP(w, r)
	log.Info("我是logHeaderHandler")
}

func setLogHeaderHandler(h http.Handler) http.Handler {
	return logHeaderHandler{h}
}

func RegisterHandlers(router *mux.Router, handlerFns ...HandlerFunc) http.Handler {
	var f http.Handler
	f = router
	for _, hFn := range handlerFns {
		f = hFn(f)
	}
	return f
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello world")
	fmt.Println(w.Header())
	fmt.Println(r.URL.Path)
	if err != nil {
		log.Error(err)
	}
}

var log = logrus.New()

func initLog(log *logrus.Logger) (*logrus.Logger, error) {
	file, err := os.OpenFile("./log/remote.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		return &logrus.Logger{}, err
	}
	log.Out = file
	return log, nil
}

func main() {
	_, err := initLog(log)
	if err != nil {
		fmt.Println(err)
	}

	remote := mux.NewRouter().PathPrefix("/api").Subrouter()

	remote.Host("s3.test.com").Methods("GET").Path("/hello").HandlerFunc(HelloWorld)

	middlewares := []HandlerFunc{
		SetCommonHeaderHandler,
		setLogHeaderHandler,
	}

	server := &http.Server{
		Addr:    ":8082",
		Handler: RegisterHandlers(remote, middlewares...),
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			err := server.ListenAndServe()
			if err != nil {
				log.Error(err)
				return
			}
		}
	}()

	wg.Wait()

}

/*
// MuxDemo4。中间件测试
func main() {
	remote := mux.NewRouter()

	remote.HandleFunc("/middle", handler)
	remote.Use(loggingMiddleware)
	remote.Use(HeaderMiddleware)

	http.ListenAndServe(":8080", remote)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Do stuff here
		fmt.Fprintln(w, "我是LoggingMiddleware")
		//fmt.Fprintf(w, "%s\r\n", r.URL)
		//fmt.Println(r.Header)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func HeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "我是HeaderMiddleware")
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("我是最终的Handler"))
	if err != nil {
		return
	}
}

*/

/*
// MuxDemo3。设置子路由、匹配前缀、域名
func main() {
	remote := mux.NewRouter()

	product := remote.Host("s3.test.com").PathPrefix("/products").Subrouter()
	product.HandleFunc("/", ProductsPage)
	product.HandleFunc("/{name}", ProductPage)

	user := remote.Host("s3.test.com").PathPrefix("/users").Subrouter()
	user.HandleFunc("/{id}/{name}", UserPage)

	http.ListenAndServe(":8080", remote)
}

func ProductsPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Products page")
	if err != nil {
		return
	}
}

func ProductPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	_, err := fmt.Fprintf(w, "ProcuceName：%s\n", vars["name"])
	if err != nil {
		return
	}
}

func UserPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	_, err := fmt.Fprintf(w, "id：%s, name：%s\n", vars["id"], vars["name"])
	if err != nil {
		return
	}
}
*/

/*
// muxDemo2。url传参，并利用正则表达式控制参数类型。
func main() {
	remote := mux.NewRouter()

	//1、传参，参数格式无限定
	// remote.HandleFunc("/intro/{title}", TitlePage)

	//2、传参，利用正则表达式限定参数格式。如下只能是字母组成的参数。输入其它类型的参数会报404错误
	remote.HandleFunc("/intro/{title:[a-z]+}", TitlePage)

	http.ListenAndServe(":8000", remote)
}

func TitlePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "title: %v\n", vars["title"])
	if err != nil {
		return
	}
}

*/

/*
//  muxDemo1，普通测试
func main() {
	remote := mux.NewRouter()
	remote.HandleFunc("/", IndexPage)
	remote.HandleFunc("/intro", IntroPage)

	http.ListenAndServe(":8000", remote)
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Hello world")
	if err != nil {
		return
	}
}

func IntroPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Introduce Page")
	if err != nil {
		return
	}
}

*/
