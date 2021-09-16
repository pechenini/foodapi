package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pechenini/foodapi/model"
	"github.com/pechenini/foodapi/server/response"
	"net/http"
	"strconv"
)

type SupplierHandler struct {
	db *model.DB
}

func NewSupplierHandler(db *model.DB) *SupplierHandler {
	return &SupplierHandler{db: db}
}

// GetSuppliers godoc
// @Summary Get suppliers
// @Description Get suppliers
// @ID get-suppliers
// @Tags Suppliers Actions
// @Produce json
// @Success 200 {object} response.SupplierCollectionResponse
// @Router /suppliers [get]
func (handler *SupplierHandler) GetSuppliers(ctx *gin.Context) {
	var suppliersResponse response.SupplierCollectionResponse
	for _, supplier := range handler.db.Suppliers {
		suppliersResponse.Suppliers = append(suppliersResponse.Suppliers, response.SupplierResponse{
			Id:    supplier.ID,
			Name:  supplier.Name,
			Type: supplier.Type,
			Image: supplier.Image,
			WorkingHours: supplier.WorkingHours,
		})
	}
	ctx.JSON(http.StatusOK, suppliersResponse)
}

// GetSupplierMenuPosition godoc
// @Summary Get supplier menu by supplier id and position id
// @Description Get menu position by supplier id and position id
// @ID get-supplier-menu-position
// @Tags Supplier Actions
// @Produce json
// @Param id path int true "Supplier ID"
// @Param position_id path int true "Supplier menu position ID"
// @Success 200 {object} model.Menu
// @Failure 400 {object} response.Error
// @Router /suppliers/{id}/menu/{position_id} [get]
func (handler *SupplierHandler) GetSupplierMenuPosition(ctx *gin.Context) {
	restId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error{
			Error: "wrong id provided",
		})
		return
	}

	positionId, err := strconv.Atoi(ctx.Param("position_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error{
			Error: "wrong position id provided",
		})
		return
	}

	for _, supplier := range handler.db.Suppliers {
		if supplier.ID != restId {
			continue
		}
		for _, menu := range supplier.Menu {
			if menu.ID == positionId {
				ctx.JSON(http.StatusOK, menu)
				return
			}
		}
	}

	ctx.JSON(http.StatusBadRequest, response.Error{
		Error: "supplier or menu position was not found",
	})
	return
}

// GetSupplierMenu godoc
// @Summary Get supplier menu by supplier id
// @Description Get menu by supplier id
// @ID get-supplier-menu
// @Tags Supplier Actions
// @Produce json
// @Param id path int true "Supplier ID"
// @Success 200 {object} response.MenuResponse
// @Failure 400 {object} response.Error
// @Router /suppliers/{id}/menu [get]
func (handler *SupplierHandler) GetSupplierMenu(ctx *gin.Context) {
	var menuResponse response.MenuResponse
	supplierIdString := ctx.Param("id")
	supplierId, err := strconv.Atoi(supplierIdString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error{
			Error: "wrong id provided",
		})
		return
	}

	for _, supplier := range handler.db.Suppliers {
		if supplier.ID == supplierId {
			menuResponse.Menu = supplier.Menu
		}
	}

	if len(menuResponse.Menu) == 0 {
		ctx.JSON(http.StatusBadRequest, response.Error{
			Error: "supplier was not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, menuResponse)
}
