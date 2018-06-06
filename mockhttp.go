package main

import (
	"flag"
	"os"
	"io/ioutil"
	"bufio"
	"bytes"
	"strings"
	"net/http"
	"fmt"
)

var apis = make(map[string]string)

var dir = flag.String("dir", "/Users/shuvigoss/.mockdir/", "mock dir")
var port = flag.String("port", "8080", "port")

func main() {
	_, err := os.Stat(*dir)
	if err != nil {
		panic("未找到工作路径 " + *dir + " " + err.Error())
	}

	infos, _ := ioutil.ReadDir(*dir)

	buildApis(infos)

	startServer()
}
func startServer() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		if strings.HasSuffix(path, "/") {
			path = path[:len(path)-1]
		}
		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		body := apis[path]
		if len(body) == 0 {
			body = "{\"status\": 404}"
		}
		writer.Write([]byte(body))
	})
	http.ListenAndServe(":" + *port, nil)
}

func buildApis(infos []os.FileInfo) {
	for _, file := range infos {
		f, _ := os.Open(*dir + file.Name())
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)
		index := 0
		var key string
		var body bytes.Buffer
		for scanner.Scan() {
			text := strings.TrimSpace(scanner.Text())
			if index == 0 {
				key = text
			} else {
				body.WriteString(text)
			}
			index ++
		}
		apis[key] = body.String()
		f.Close()
	}

	fmt.Println("Registed Apis ")
	for k, _ := range apis {
		fmt.Println(k)
	}

}
