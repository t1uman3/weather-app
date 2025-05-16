package model

type Favorite struct {
	ID   int    `json:"id"`
	City string `json:"city"`
}

type FavoritesList struct {
	Favorites []Favorite `json:"favorites,omitempty"`
}
