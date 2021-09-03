package response

import "github.com/pechenini/foodapi/model"

type RestaurantCollectionResponse struct {
	Restaurants []RestaurantResponse `json:"restaurants"`
}

type RestaurantResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Image string `json:"image"`
}

type MenuResponse struct {
	Menu []model.Menu `json:"menu"`
}


type Error struct {
	Error string `json:"error"`
}