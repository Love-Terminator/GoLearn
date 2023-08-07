package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// MuxDemo4。中间件测试
func main() {
	remote := mux.NewRouter()

	remote.HandleFunc("/middle", handler)
	remote.Use(loggingMiddleware)

	http.ListenAndServe(":8080", remote)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Do stuff here
		fmt.Println(r.RequestURI)
		fmt.Fprintf(w, "%s\r\n", r.URL)
		fmt.Println(r.Header)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("handle middleware"))
	fmt.Println("print handler")
}

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
