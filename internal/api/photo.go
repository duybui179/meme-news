package api

import (
	"encoding/json"
	"net/http"
	"dev.duybt/internal/models"
)

func GetPhotoByURL(url string) ([]models.Photo, error){
	return fetchPhotos(url)
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