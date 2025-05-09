package service

import (
	"encoding/json"
	"os"
	"sync"
)

type FavoriteService struct {
	favorites map[string]bool
	mutex     sync.RWMutex
	filePath  string
}

// NewFavoriteService создает сервис и загружает избранные из файла
func NewFavoriteService(filePath string) *FavoriteService {
	s := &FavoriteService{
		favorites: make(map[string]bool),
		filePath:  filePath,
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

	var cities []string
	if err := json.NewDecoder(file).Decode(&cities); err == nil {
		for _, city := range cities {
			s.favorites[city] = true
		}
	}
}

// saveToFile сохраняет избранные города в файл
func (s *FavoriteService) saveToFile() {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	cities := make([]string, 0, len(s.favorites))
	for city := range s.favorites {
		cities = append(cities, city)
	}

	file, err := os.Create(s.filePath)
	if err != nil {
		// Можно добавить логирование ошибки
		return
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	_ = enc.Encode(cities)
}

func (s *FavoriteService) AddFavorite(city string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.favorites[city] = true
	go s.saveToFile()
}

func (s *FavoriteService) RemoveFavorite(city string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.favorites, city)
	go s.saveToFile()
}

func (s *FavoriteService) GetFavorites() []string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	cities := make([]string, 0, len(s.favorites))
	for city := range s.favorites {
		cities = append(cities, city)
	}
	return cities
}
