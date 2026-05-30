package api

import (
	"encoding/json"
	"net/http"
	"dev.duybt/internal/models"
	"fmt"
)

func GetPhotoByURL(url string) ([]models.Photo, error){
	return fetchPhotos(url)
}

type result struct {
    photos []models.Photo
    err    error
}

func GetContenGoroutine(urls []string) [][]models.Photo {
	ch := make(chan result)

	for _, u := range urls {
		go func(url string) {
			photos, err := fetchPhotos(url)
			ch <- result{photos: photos, err: err}
		}(u)
	}

	var all [][]models.Photo 
	for range urls {
		res := <- ch 
		 if res.err != nil {
            fmt.Println("Error:", res.err)
            continue
        }
        all = append(all, res.photos)
    }
    return all
}

func fetchPhotos(url string) ([]models.Photo, error){
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var photos []models.Photo
	if err := json.NewDecoder(resp.Body).Decode(&photos);  err != nil {
		return nil, err
	}
	return photos, nil
}