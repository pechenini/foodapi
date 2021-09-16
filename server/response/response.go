package response

import "github.com/pechenini/foodapi/model"

type SupplierCollectionResponse struct {
	Suppliers []SupplierResponse `json:"restaurants"`
}

type SupplierResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Image string `json:"image"`
	WorkingHours model.WorkingHours `json:"workingHours"`
}

type MenuResponse struct {
	Menu []model.Menu `json:"menu"`
}


type Error struct {
	Error string `json:"error"`
}