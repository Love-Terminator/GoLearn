package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	testRouter := router.Host("s3.test.com").PathPrefix("/").Subrouter()
	testRouter.Methods("POST").Path("/form").HandlerFunc(xWwwFormUrlencoded)
	testRouter.Methods("GET").Path("/get").HandlerFunc(getTest)

	server := &http.Server{
		Addr:    ":8100",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Lister Error!")
	}
}

// content-type: x-www-form-urlencoded
// 将请求参数放在Request body中，如果是中文或特殊字符如"/"、","、“:" 等会自动进行URL转码。
// 不支持文件，一般用于表单提交，即POST请求方式
func xWwwFormUrlencoded(w http.ResponseWriter, r *http.Request) {
	// 获取body的长度
	bodyLen := r.ContentLength
	// 字节切片
	body := make([]byte, bodyLen)
	// 获取request body
	r.Body.Read(body)
	// 输出request body
	fmt.Fprintln(w, "body: "+string(body))

}

func getTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "getTest")
}
