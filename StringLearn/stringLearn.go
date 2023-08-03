package main

import (
	"fmt"
	"gopkg.in/bufio.v1"
	"io"
)

var TableSeparator string = string([]byte{92})
var TableObjectPrefix string = string([]byte{0xFF}) + "yig" + "o"
var TableMinKeySuffix string = ""
var TableMaxKeySuffix string = string([]byte{0xFF})

func GenKey(args ...string) []byte {
	buf := bufio.NewBuffer([]byte{})
	for _, arg := range args {
		buf.WriteString(arg)
		buf.WriteString(TableSeparator)
	}
	key := buf.Bytes()

	return key[:len(key)-1]
}

type ThrottleWriter struct {
	writer io.Writer
	t      *Throttler
}

func (t ThrottleWriter) Write(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

type Throttler struct {
	UserID        string
	AcquiredToken int64
	Limit         int64
}

func NewThrottleWriter(writer io.Writer) *ThrottleWriter {
	return &ThrottleWriter{
		writer: writer,
		t: &Throttler{
			UserID: "123",
			Limit:  565,
		},
	}
}

func main() {
	bucketName := "ypc"
	objectName := "ceph.pdf"
	fmt.Println(GenKey(TableObjectPrefix, bucketName, objectName, TableMinKeySuffix))
	fmt.Println(GenKey(TableObjectPrefix, bucketName, objectName, TableMaxKeySuffix))
	fmt.Println(TableMaxKeySuffix)

	var writer io.Writer
	writer = NewThrottleWriter(writer)
}
