package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/t1uman3/weather-app/backend-go/api"
	"github.com/t1uman3/weather-app/backend-go/service"
)

func main() {
	path := ".env"
	for {
		err := godotenv.Load(path)
		if err == nil {
			break
		}
		path = "../" + path
	}

	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		log.Println("Предупреждение: WEATHER_API_KEY не установлен")
	}

	weatherService := service.NewWeatherService(apiKey)
	weatherHandler := api.NewWeatherHandler(weatherService)

	favoriteService := service.NewFavoriteService("favorites.json")
	favoriteHandler := api.NewFavoriteHandler(favoriteService)

	e := echo.New()

	// Настройка middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/weather", weatherHandler.HandleGetWeather)
	e.POST("/weather", weatherHandler.HandlePostWeather)
	e.GET("/favorites", favoriteHandler.HandleGetFavorites)
	e.POST("/favorites", favoriteHandler.HandleAddFavorite)
	e.DELETE("/favorites/:city", favoriteHandler.HandleRemoveFavorite)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Запуск сервера
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
