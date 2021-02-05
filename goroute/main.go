package main

import (
	"log"
	"net/http"
	"time"
)

func task()  {
	var v int = 0
	for i := 0; i < 5000; i++ {
		time.Sleep(1*time.Millisecond)
		v = i * v
	}
	_ = v
}

func handler(w http.ResponseWriter, r *http.Request) {
	go task()
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/", handler) //	设置访问路由
	log.Fatal(http.ListenAndServe(":7080", nil))
}
