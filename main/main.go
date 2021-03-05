package main

import (
	"html/template"
	"log"
	"net/http"
)

// 用于测试代码
// like
// import _ "package path"
func main() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	files, _ := template.ParseFiles("test.html")
	files.Execute(w, nil)
}
