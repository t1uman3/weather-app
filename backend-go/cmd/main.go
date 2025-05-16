package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/t1uman3/weather-app/backend-go/api"
	"github.com/t1uman3/weather-app/backend-go/service"
)

func main() {
	fmt.Println("Starting backend service...")

	// Попытаемся загрузить .env файл, но не будем выходить с ошибкой, если его нет
	// Вместо этого используем переменные окружения из Docker
	_ = godotenv.Load(".env")
	fmt.Println("Environment check completed")

	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		log.Println("Предупреждение: WEATHER_API_KEY не установлен")
	} else {
		fmt.Println("WEATHER_API_KEY loaded successfully")
	}

	weatherService := service.NewWeatherService(apiKey)
	weatherHandler := api.NewWeatherHandler(weatherService)
	fmt.Println("Weather service initialized")

	favoriteService := service.NewFavoriteService("favorites.json")
	favoriteHandler := api.NewFavoriteHandler(favoriteService)
	fmt.Println("Favorite service initialized")

	e := echo.New()
	e.HideBanner = false
	e.Debug = true
	fmt.Println("Echo server initialized")

	// Настройка middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	fmt.Println("Middleware configured")

	// Маршруты API
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": "backend is ready",
			"version": "0.1.0",
		})
	})
	fmt.Println("Route / registered")

	// Маршруты погоды - совместимые с Python-бэкендом
	e.POST("/weather", weatherHandler.HandlePostWeather)
	fmt.Println("Route /weather registered")

	// Маршруты избранных городов - совместимые с Python-бэкендом
	e.GET("/favorites", favoriteHandler.HandleGetFavorites)
	e.POST("/favorite", favoriteHandler.HandleAddFavorite)
	e.DELETE("/favorite/:city", favoriteHandler.HandleRemoveFavorite)
	fmt.Println("Favorite routes registered")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Запуск сервера
	address := "0.0.0.0:" + port
	fmt.Printf("Starting server on %s...\n", address)
	if err := e.Start(address); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
