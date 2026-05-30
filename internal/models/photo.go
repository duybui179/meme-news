package models

type Photo struct {
	AlbumId int `json:"albumId"`
	ID int `json:"id"`
	Title string `json:"title"`
	URL string `json:"url"`
	ThumbnailURL string `json:"thumbnail"`
}