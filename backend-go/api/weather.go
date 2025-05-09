package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/t1uman3/weather-app/backend-go/model"
	"github.com/t1uman3/weather-app/backend-go/service"
)

type WeatherHandler struct {
	weatherService *service.WeatherService
}

func NewWeatherHandler(weatherService *service.WeatherService) *WeatherHandler {
	return &WeatherHandler{
		weatherService: weatherService,
	}
}

func (h *WeatherHandler) HandleGetWeather(c echo.Context) error {
	// Получаем город из запроса
	city := c.QueryParam("city")
	if city == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "Параметр 'city' обязателен",
		})
	}

	weather, err := h.weatherService.GetWeatherByCity(city)
	if err != nil {
		// Логирование ошибки
		c.Logger().Error(err)

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: "Ошибка при получении погоды: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, weather)
}

func (h *WeatherHandler) HandlePostWeather(c echo.Context) error {
	// Структура для разбора входящего JSON
	var request struct {
		City string `json:"city"`
	}

	// Разбор JSON из тела запроса
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "Неверный формат запроса",
		})
	}

	// Проверка наличия города
	if request.City == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "Параметр 'city' обязателен",
		})
	}

	// Получаем погоду через сервис
	weather, err := h.weatherService.GetWeatherByCity(request.City)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: "Ошибка при получении погоды: " + err.Error(),
		})
	}

	// Отправляем успешный ответ
	return c.JSON(http.StatusOK, weather)
}
