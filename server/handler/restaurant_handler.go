package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pechenini/foodapi/model"
	"github.com/pechenini/foodapi/server/response"
	"net/http"
	"strconv"
)

type RestaurantHandler struct {
	db model.DB
}

func NewRestaurantHandler(db model.DB) *RestaurantHandler {
	return &RestaurantHandler{db: db}
}

// GetRestaurants godoc
// @Summary Get restaurants
// @Description Get restaurants
// @ID get-restaurants
// @Tags Restaurants Actions
// @Produce json
// @Success 200 {object} response.RestaurantCollectionResponse
// @Router /restaurants [get]
func (handler *RestaurantHandler) GetRestaurants(ctx *gin.Context) {
	var restaurantsResponse response.RestaurantCollectionResponse
	for _, restaurant := range handler.db.Restaurants {
		restaurantsResponse.Restaurants = append(restaurantsResponse.Restaurants, response.RestaurantResponse{
			Id:   restaurant.ID,
			Name: restaurant.Name,
		})
	}
	ctx.JSON(http.StatusOK, restaurantsResponse)
}

// GetRestaurantMenu godoc
// @Summary Get restaurant menu by restaurant id
// @Description Get menu by restaurant id
// @ID get-restaurant-menu
// @Tags Restaurant Actions
// @Produce json
// @Param id path int true "Restaurant ID"
// @Success 200 {object} []model.Menu
// @Failure 400 {object} response.Error
// @Router /restaurants/{id}/menu [get]
func (handler *RestaurantHandler) GetRestaurantMenu(ctx *gin.Context) {
	var menuResponse response.MenuResponse
	restIdString := ctx.Param("id")
	restId, err := strconv.Atoi(restIdString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error{
			Error: "wrong id provided",
		})
		return
	}

	for _, restaurant := range handler.db.Restaurants {
		if restaurant.ID == restId {
			menuResponse.Menu = restaurant.Menu
		}
	}

	if len(menuResponse.Menu) == 0 {
		ctx.JSON(http.StatusBadRequest, response.Error{
			Error: "restaurant was not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, menuResponse)
}
