package service

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/t1uman3/weather-app/backend-go/model"
)

type favoriteData struct {
	Favorites []model.Favorite `json:"favorites"`
	LastID    int              `json:"last_id"`
}

type FavoriteService struct {
	data     favoriteData
	mutex    sync.RWMutex
	filePath string
}

// NewFavoriteService создает сервис и загружает избранные из файла
func NewFavoriteService(filePath string) *FavoriteService {
	s := &FavoriteService{
		data: favoriteData{
			Favorites: []model.Favorite{},
			LastID:    0,
		},
		filePath: filePath,
	}
	s.loadFromFile()
	return s
}

// loadFromFile загружает избранные города из файла
func (s *FavoriteService) loadFromFile() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	file, err := os.Open(s.filePath)
	if err != nil {
		// Если файла нет - ничего страшного, просто создадим новый позже
		return
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&s.data); err != nil {
		// Если ошибка чтения, инициализируем пустыми данными
		s.data = favoriteData{
			Favorites: []model.Favorite{},
			LastID:    0,
		}
	}
}

// saveToFile сохраняет избранные города в файл
func (s *FavoriteService) saveToFile() {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	file, err := os.Create(s.filePath)
	if err != nil {
		// Можно добавить логирование ошибки
		return
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	_ = enc.Encode(s.data)
}

func (s *FavoriteService) AddFavorite(city string) model.Favorite {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Проверяем, есть ли уже такой город
	for _, fav := range s.data.Favorites {
		if fav.City == city {
			return fav
		}
	}

	// Добавляем новый город
	s.data.LastID++
	newFavorite := model.Favorite{
		ID:   s.data.LastID,
		City: city,
	}
	s.data.Favorites = append(s.data.Favorites, newFavorite)
	go s.saveToFile()

	return newFavorite
}

func (s *FavoriteService) RemoveFavorite(cityID int) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, fav := range s.data.Favorites {
		if fav.ID == cityID {
			// Удаляем элемент из слайса
			s.data.Favorites = append(s.data.Favorites[:i], s.data.Favorites[i+1:]...)
			go s.saveToFile()
			return true
		}
	}

	return false
}

func (s *FavoriteService) GetFavorites() []model.Favorite {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Возвращаем копию слайса, чтобы избежать гонок данных
	result := make([]model.Favorite, len(s.data.Favorites))
	copy(result, s.data.Favorites)
	return result
}
