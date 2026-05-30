package api

import(
	"fmt"
	"dev.duybt/internal/api"
)

func GetContent(){
	result := "meme"
	fmt.Println("Check content", result)

	api.GetContent()
	
	url := "https://jsonplaceholder.typicode.com/photos"

	photos, err := api.GetPhotoByURL(url)

	if err != nil {
        panic(err)
    }

	 for i := 0; i < 5; i++ {
        fmt.Printf("%d: %s (%s)\n", photos[i].ID, photos[i].Title, photos[i].URL)
    }

	 urls := []string{
        "https://jsonplaceholder.typicode.com/photos",
        "https://jsonplaceholder.typicode.com/comments", // ví dụ thêm
    }

	results := api.GetContenGoroutine(urls)

	for i, photos := range results {
        fmt.Printf("URL %d có %d phần tử\n", i+1, len(photos))
    }
}