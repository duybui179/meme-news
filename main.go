package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chào bạn! Đây là Go App chạy trên port 8081 🚀")
}

func main() {

	http.HandleFunc("/", helloHandler)

	port := ":8081"
	fmt.Printf("Server đang khởi động tại http://localhost%s\n", port)
	
	// Bắt đầu lắng nghe và phục vụ
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Lỗi khởi động server: %s\n", err)
	}
}