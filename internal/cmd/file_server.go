package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	baseDir = "/uploads"
)

// /static && /upload
func serveFile() {
	fs := http.FileServer(http.Dir(baseDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/upload", fileUploadHandler)
	http.ListenAndServe(":9090", nil)
}

func fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	//设置内存大小
	r.ParseMultipartForm(32 << 20)
	//创建上传目录
	paths := r.URL.Query()["paths"]
	for i := range paths {
		os.Mkdir(baseDir+"/"+strings.Join(paths[0:i+1], "/"), os.ModePerm)
	}
	//获取上传的文件组
	files := r.MultipartForm.File["uploadfile"]
	len := len(files)
	for i := 0; i < len; i++ {
		//打开上传文件
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			log.Errorln(err)
		}

		//创建上传文件
		cur, err := os.Create(baseDir + "/" + strings.Join(paths, "/") + "/" + files[i].Filename)
		defer cur.Close()
		if err != nil {
			log.Errorln(err)
		}
		io.Copy(cur, file)
		log.Infoln(files[i].Filename, "uploaded")
	}
	//输出上传的文件名
	fmt.Fprintf(w, "success!\n") //这个写入到w的是输出到客户端的
}
