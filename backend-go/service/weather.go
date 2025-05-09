package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/t1uman3/weather-app/backend-go/model"
)

// WeatherService предоставляет методы для работы с погодными данными
type WeatherService struct {
	apiKey string
}

// NewWeatherService создает новый сервис погоды
func NewWeatherService(apiKey string) *WeatherService {
	return &WeatherService{
		apiKey: apiKey,
	}
}

// GetWeatherByCity получает погоду по названию города
func (s *WeatherService) GetWeatherByCity(city string) (model.Weather, error) {
	// Базовая валидация
	if strings.TrimSpace(city) == "" {
		return model.Weather{}, fmt.Errorf("название города не может быть пустым")
	}

	// Проверка ключа API
	if s.apiKey == "" {
		return model.Weather{}, fmt.Errorf("API ключ не установлен")
	}

	// Создаем URL с параметрами
	baseURL := "https://api.openweathermap.org/data/2.5/weather"
	params := url.Values{}
	params.Add("q", city)
	params.Add("appid", s.apiKey)
	params.Add("units", "metric") // Температура в Цельсиях
	params.Add("lang", "ru")      // Локализация на русском

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Выполняем запрос
	resp, err := http.Get(fullURL)
	if err != nil {
		return model.Weather{}, fmt.Errorf("ошибка HTTP-запроса: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return model.Weather{}, fmt.Errorf("API вернул статус: %d", resp.StatusCode)
	}

	// Парсим ответ
	var openWeather struct {
		Name string `json:"name"`
		Main struct {
			Temp     float64 `json:"temp"`
			Humidity int     `json:"humidity"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Wind struct {
			Speed float64 `json:"speed"`
		} `json:"wind"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&openWeather); err != nil {
		return model.Weather{}, fmt.Errorf("ошибка при разборе JSON: %w", err)
	}

	// Формируем и возвращаем данные
	weather := model.Weather{
		City:        openWeather.Name,
		Temperature: openWeather.Main.Temp,
		Humidity:    openWeather.Main.Humidity,
		WindSpeed:   openWeather.Wind.Speed,
	}

	if len(openWeather.Weather) > 0 {
		weather.Description = openWeather.Weather[0].Description
		weather.Icon = openWeather.Weather[0].Icon
	}

	return weather, nil
}
