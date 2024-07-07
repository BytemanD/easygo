package http

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/BytemanD/easygo/pkg/global/logging"
	"github.com/BytemanD/easygo/pkg/sysutils"
	"github.com/gin-gonic/gin"
)

var HTML = `
<!DOCTYPE html>
<html lang="zh-CN">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<link rel="stylesheet" href="https://cdn.staticfile.org/font-awesome/4.7.0/css/font-awesome.css">
		<title>SimpleHttpFS</title>
	</head>
	<body>
		<div style="margin-left: 50px">
			<table>
				<tr></th><th>类型</th><th>名称</th><th>大小</th></tr>
				{{ range $index, $webFile := . }}
					<tr>
						{{ if $webFile.IsDir }}
						<td><i class="fa fa-folder" style="color: orange;"></i></td>
						<td><a href="{{$webFile.WebPath}}" > {{ $webFile.Name }} </a></li></td>
						{{ else }}
						<td><i class="fa fa-file" style="color: grey;"></i></td>
						<td><a target="view_window" href="{{$webFile.WebPath}}" > {{ $webFile.Name }} </a></li></td>
						{{ end }}
						<td style="text-align:right">{{ $webFile.HumanSize }}</td>
					</tr>
				{{ end}}
			<table>
		</div>
</body>
<html>
`
var FSConfig HTTPFSConfig

type HTTPFSConfig struct {
	Port int16
	Root string
}

type WebFile struct {
	Dir  string
	Name string
	Size int64
}

const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
)

func (webFile *WebFile) LogicPath() string {
	return filepath.Join(webFile.Dir, webFile.Name)
}
func (webFile *WebFile) WebPath() string {
	return strings.ReplaceAll(webFile.LogicPath(), "\\", "/")
}

func (webFile *WebFile) IsDir() bool {
	filePath := filepath.Join(FSConfig.Root, webFile.LogicPath())
	fi, _ := os.Stat(filePath)
	return fi.IsDir()
}
func (webFile *WebFile) HumanSize() string {
	switch {
	case webFile.Size > GB:
		return fmt.Sprintf("%.1f G", float32(webFile.Size)/GB)
	case webFile.Size > MB:
		return fmt.Sprintf("%.1f M", float32(webFile.Size)/MB)
	case webFile.Size > KB:
		return fmt.Sprintf("%.1f K", float32(webFile.Size)/KB)
	default:
		return fmt.Sprintf("%d B", webFile.Size)
	}
}

func handleFileDownload(respWriter http.ResponseWriter, request *http.Request) {
	filePath := filepath.Join(FSConfig.Root, request.URL.Path)
	filePath = strings.ReplaceAll(filePath, "\\", "/")
	file, _ := os.Open(filePath)
	defer file.Close()

	logging.Info("下载文件: %s", file.Name())
	respWriter.Header().Set("Content-Disposition", "attachment; filename="+
		filepath.Base(file.Name()))
	http.ServeFile(respWriter, request, filePath)
}

type VuetifyHttpFS struct {
	Port          int16
	Root          string
	StaticPath    string
	staticFileMap string
}

type DirEntry struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
	Size  int64  `json:"size"`
}

func getDirEntries(dirPath string) ([]DirEntry, error) {
	entries := []DirEntry{}
	dirEnties, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	for _, entry := range dirEnties {
		var size int64
		if !entry.IsDir() {
			if stat, err := os.Stat(path.Join(dirPath, entry.Name())); err == nil {
				size = stat.Size()
			}
		}
		entries = append(entries, DirEntry{
			Name:  entry.Name(),
			IsDir: entry.IsDir(),
			Size:  size,
		})
	}
	return entries, nil
}
func (s VuetifyHttpFS) root(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "index.html")
}
func (s VuetifyHttpFS) index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "VuetifyHttpFS",
	})
}

func (s VuetifyHttpFS) getQueryPath(c *gin.Context) string {
	dirPath := c.Query("path")
	if strings.HasPrefix(dirPath, "/") {
		dirPath = dirPath[1:]
	}
	return dirPath
}
func (s VuetifyHttpFS) getEntries(c *gin.Context) {
	dirPath := s.getQueryPath(c)
	entries, err := getDirEntries(path.Join(s.Root, dirPath))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"entries": entries})
}

func (s VuetifyHttpFS) uploadFile(c *gin.Context) {
	dirPath := s.getQueryPath(c)
	saveDir := path.Join(s.Root, dirPath)
	form, _ := c.MultipartForm()
	files := form.File["file"]
	var err error
	for _, file := range files {
		saveFile := path.Join(saveDir, file.Filename)
		logging.Info("saving file to %s", saveFile)
		if err := c.SaveUploadedFile(file, saveFile); err != nil {
			break
		}
		logging.Info("saved file %s", saveFile)
	}
	if err != nil {
		c.JSON(200, gin.H{
			"result": "上传失败",
			"error":  err.Error(),
		})
	}
}
func (s VuetifyHttpFS) deleteFile(c *gin.Context) {
	dirPath := s.getQueryPath(c)
	filePath := path.Join(s.Root, dirPath)
	if stat, err := os.Stat(filePath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		if stat.IsDir() {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s is dir", dirPath)})
		} else {
			os.Remove(filePath)
			c.JSON(http.StatusNoContent, gin.H{})
		}
	}
}
func (s VuetifyHttpFS) getServerAddr() string {
	return fmt.Sprintf(":%d", s.Port)
}
func (s VuetifyHttpFS) Run() error {
	r := gin.Default()
	r.Delims("[[", "]]")
	r.LoadHTMLGlob(s.StaticPath + "/*")

	r.GET("/", s.root)
	r.GET("/index.html", s.index)

	r.GET("/fs/entries", s.getEntries)
	r.POST("/fs/entries", s.uploadFile)
	r.DELETE("/fs/entries", s.deleteFile)

	ipaddrs, err := sysutils.GetAllIpaddress()
	if err != nil {
		return err
	}
	webAddr := []string{}
	for _, ipaddr := range ipaddrs {
		webAddr = append(webAddr, fmt.Sprintf("http://%s:%d", ipaddr, s.Port))
	}
	logging.Info("启动web服务:\n----\n%s\n----", strings.Join(webAddr, "\n"))

	return r.Run(fmt.Sprintf(s.getServerAddr()))
}
