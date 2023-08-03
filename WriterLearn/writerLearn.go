package main

import (
	"fmt"
)

// HttpRange specifies the byte range to be sent to the client.
type HttpRange struct {
	OffsetBegin  int64
	OffsetEnd    int64
	ResourceSize int64
}

type GetObjectResponseWriter struct {
	statusCode int
	version    string
}

func (g GetObjectResponseWriter) Write(p []byte) (n int, err error) {
	//TODO implement me
	fmt.Println(g)

	return 3, nil
}

func newGetObjectResponseWriter(statusCode int, version string) *GetObjectResponseWriter {
	return &GetObjectResponseWriter{statusCode, version}
}

func main() {

	statusCode := 404
	version := "lastest"

	p := []byte{'y', 'i', 'g'}
	write, err := newGetObjectResponseWriter(statusCode, version).Write(p)
	if err != nil {
		return
	}
	fmt.Println(write)

	type Bytes []byte
	b := Bytes{232, 191, 153, 230, 152, 175, 228, 184, 128, 228, 184, 170, 229, 173, 151, 231, 172, 166, 228, 184, 178}
	s := string(b)
	fmt.Println(s)

	fmt.Println(string([]byte{92}))

}
