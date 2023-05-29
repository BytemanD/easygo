package http

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var HTML = `
<!DOCTYPE html>
<html lang="zh-CN">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=devie-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>SimpleHttpServer</title>
	</head>
	<body>
	<ol>
	{{ range $index, $dirName := . }}
		<li> <a href="{{ $dirName}}"> {{ $dirName}} </a></li>
	{{ end}}
	</ol>
</body>
<html>
`

func FilePathHandler(respWriter http.ResponseWriter, request *http.Request) {
	log.Printf("请求地址 %s", request.URL.Path)
	dirPath := fmt.Sprintf(".%s", request.URL.Path)
	rd, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Fprintf(respWriter, fmt.Sprintf("非法路径 %s", dirPath))
		return
	}
	dirList := []string{}
	for _, fi := range rd {
		dirList = append(dirList, fi.Name())
	}
	tmpl, err := template.New("test").Parse(HTML)
	tmpl.Execute(respWriter, dirList)
}

func SimpleHttpServer(port int16) {
	http.HandleFunc("/", FilePathHandler) //设置访问的路由

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("启动HTTP服务失败: ", err)
	}

}
