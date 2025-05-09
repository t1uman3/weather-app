package model

type Favorite struct {
	City string `json:"city"`
}

type FavoritesList struct {
	Favorites []string `json:"favorites"`
}
