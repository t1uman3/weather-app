package api

import (
	"net/http"

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
	return c.JSON(http.StatusOK, model.FavoritesList{
		Favorites: favorites,
	})
}

// HandleAddFavorite добавляет город в избранное
func (h *FavoriteHandler) HandleAddFavorite(c echo.Context) error {
	var favorite model.Favorite
	if err := c.Bind(&favorite); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "Неверный формат запроса",
		})
	}

	if favorite.City == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "Название города не может быть пустым",
		})
	}

	h.favoriteService.AddFavorite(favorite.City)
	return c.JSON(http.StatusOK, map[string]bool{"success": true})
}

// HandleRemoveFavorite удаляет город из избранного
func (h *FavoriteHandler) HandleRemoveFavorite(c echo.Context) error {
	city := c.Param("city")
	if city == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "Название города не может быть пустым",
		})
	}

	h.favoriteService.RemoveFavorite(city)
	return c.JSON(http.StatusOK, map[string]bool{"success": true})
}
