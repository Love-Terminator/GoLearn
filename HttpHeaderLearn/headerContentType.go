package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Info struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func main() {
	router := mux.NewRouter()

	testRouter := router.Host("s3.test.com").PathPrefix("/").Subrouter()
	testRouter.Methods("POST").Path("/form").HandlerFunc(xWwwFormUrlencoded)
	testRouter.Methods("POST").Path("/formdata").HandlerFunc(formData)
	testRouter.Methods("POST").Path("/formstream").HandlerFunc(OctetStream)

	server := &http.Server{
		Addr:    ":8100",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Lister Error!")
	}
}

// 一、
// content-type: application/x-www-form-urlencoded
// 将请求参数放在Request body中，如果是中文或特殊字符如"/"、","、“:" 等会自动进行URL转码。
// 不支持文件，一般用于表单提交，即POST请求方式
func xWwwFormUrlencoded(w http.ResponseWriter, r *http.Request) {

	/*
		// 一、解析request
		log.Printf("Get %s request. host: %s\n", r.Method, r.Host)
		// 获取body的长度
		contentLen := r.ContentLength
		rtype := r.Header.Get("Content-Type")
		// 定义字节切片
		content := make([]byte, contentLen)
		// 获取request body
		r.Body.Read(content)
		fmt.Fprintln(w, "request-type: "+rtype)
		fmt.Fprintln(w, "request-lenght: "+strconv.FormatInt(contentLen, 10))
		fmt.Fprintln(w, "body: "+string(content))
	*/

	// 二、解析body为map类型。针对application/x-www-form-urlencoded类型
	r.ParseForm()
	params := r.Form
	fmt.Fprintln(w, params)
}

// 二、
// content-type：multipart/form-data
// 首先生成了一个 boundary 用于分割不同的字段，在请求实体里每个参数以------boundary开始，然后是附加信息和参数名，
// 然后是空行，最后是参数内容。多个参数将会有多个boundary块。如果参数是文件会有特别的文件域。最后以------boundary–为结束标识。
// multipart/form-data支持文件上传的格式，一般需要上传文件的表单则用该类型。
func formData(w http.ResponseWriter, r *http.Request) {

	/*
		// 一、解析request
		log.Printf("Get %s request. host: %s\n", r.Method, r.Host)
		// 获取body的长度
		contentLen := r.ContentLength
		rtype := r.Header.Get("Content-Type")
		// 定义字节切片
		content := make([]byte, contentLen)
		// 获取request body
		r.Body.Read(content)
		fmt.Fprintln(w, "request-type: "+rtype)
		fmt.Fprintln(w, "request-lenght: "+strconv.FormatInt(contentLen, 10))
		fmt.Fprintln(w, "body: "+string(content))

	*/

	// 一、解析body为map类型。针对multipart/form-data类型
	r.ParseMultipartForm(128)
	params := r.MultipartForm

	// 获取普通类型的参数
	values := params.Value
	for key, value := range values {
		fmt.Fprint(w, key+": ")
		fmt.Fprintln(w, value)
	}

	// 获取file类型的参数
	picFiles := params.File
	fmt.Fprintln(w, picFiles)
	for _, files := range picFiles {
		for i := 0; i < len(files); i++ {
			fileHeader, _ := files[i].Open()

			fileBuffer := make([]byte, files[i].Size)
			_, err := fileHeader.Read(fileBuffer)
			if err != nil {
				log.Fatalln("读取文件数据有误")
				return
			}

			copyFile, _ := os.OpenFile("./copyFile/"+files[i].Filename, os.O_CREATE|os.O_RDWR, 0777)
			_, err = copyFile.Write(fileBuffer)
			if err != nil {
				log.Fatalln()
				return
			}
			defer copyFile.Close()
		}
	}
}

func OctetStream(w http.ResponseWriter, r *http.Request) {
	/*
		stream := make([]byte, r.ContentLength)
		reader := bufio.NewReader(r.Body)
		_, err := reader.Read(stream)
		if err != nil {
			return
		}
		fmt.Fprintln(w, stream) // 二进制数据

	*/

	buff := make([]byte, r.ContentLength)
	r.Body.Read(buff)
	fmt.Fprintln(w, buff)
}
