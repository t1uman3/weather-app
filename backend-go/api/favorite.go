package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/t1uman3/weather-app/backend-go/model"
	"github.com/t1uman3/weather-app/backend-go/service"
)

type FavoriteHandler struct {
	favoriteService *service.FavoriteService
}

// NewFavoriteHandler создает новый обработчик избранных городов
func NewFavoriteHandler(favoriteService *service.FavoriteService) *FavoriteHandler {
	return &FavoriteHandler{
		favoriteService: favoriteService,
	}
}

// HandleGetFavorites возвращает список избранных городов
func (h *FavoriteHandler) HandleGetFavorites(c echo.Context) error {
	favorites := h.favoriteService.GetFavorites()
	return c.JSON(http.StatusOK, favorites)
}

// HandleAddFavorite добавляет город в избранное
func (h *FavoriteHandler) HandleAddFavorite(c echo.Context) error {
	var request struct {
		City string `json:"city"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "Неверный формат запроса",
		})
	}

	if request.City == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "Название города не может быть пустым",
		})
	}

	favorite := h.favoriteService.AddFavorite(request.City)
	return c.JSON(http.StatusOK, favorite)
}

// HandleRemoveFavorite удаляет город из избранного
func (h *FavoriteHandler) HandleRemoveFavorite(c echo.Context) error {
	cityIDStr := c.Param("city")
	cityID, err := strconv.Atoi(cityIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "Некорректный ID города",
		})
	}

	success := h.favoriteService.RemoveFavorite(cityID)
	if !success {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			Error: "Город не найден",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"detail": "City deleted",
	})
}
