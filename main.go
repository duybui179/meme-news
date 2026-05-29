package main

import (
	"fmt"
	"net/http"
	"dev.duybt/internal/api"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chào bạn! Đây là Go App chạy trên port 8081 🚀")
}

func main() {
	api.GetContent()
	
	url := "https://jsonplaceholder.typicode.com/photos"

	photos, err := api.GetPhotoByURL(url)

	if err != nil {
        panic(err)
    }

	 for i := 0; i < 5; i++ {
        fmt.Printf("%d: %s (%s)\n", photos[i].ID, photos[i].Title, photos[i].URL)
    }
	
	// http.HandleFunc("/", helloHandler)

	// port := ":8081"
	// fmt.Printf("Server đang khởi động tại http://localhost%s\n", port)
	
	// // Bắt đầu lắng nghe và phục vụ
	// if err := http.ListenAndServe(port, nil); err != nil {
	// 	fmt.Printf("Lỗi khởi động server: %s\n", err)
	// }
}