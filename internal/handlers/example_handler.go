package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/myorg/myservice/internal/models"
	"github.com/myorg/myservice/internal/services"
	"github.com/myorg/myservice/pkg/response"
)

// ItemHandler struct holding dependencies
type ItemHandler struct {
	service services.ItemService
}

// NewItemHandler factory
func NewItemHandler(service services.ItemService) *ItemHandler {
	return &ItemHandler{service}
}

// GetItems handles GET /items
// @Summary      Get all items
// @Description  Get a list of all items
// @Tags         items
// @Produce      json
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.ErrorResponse
// @Router       /items [get]
func (h *ItemHandler) GetItems(c *gin.Context) {
	items, err := h.service.GetAllItems()
	if err != nil {
		response.InternalServerError(c, err)
		return
	}
	response.Success(c, http.StatusOK, "Items fetched successfully", items)
}

// GetItem handles GET /items/:id
// @Summary      Get an item by ID
// @Description  Get a single item by its ID
// @Tags         items
// @Produce      json
// @Param        id   path      int  true  "Item ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Router       /items/{id} [get]
func (h *ItemHandler) GetItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	item, err := h.service.GetItemByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Item not found", err)
		return
	}

	response.Success(c, http.StatusOK, "Item fetched successfully", item)
}

// CreateItem handles POST /items
// @Summary      Create a new item
// @Description  Create a new item with the input payload
// @Tags         items
// @Accept       json
// @Produce      json
// @Param        item body      models.Item  true  "Item payload"
// @Success      201  {object}  response.Response
// @Failure      400  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /items [post]
func (h *ItemHandler) CreateItem(c *gin.Context) {
	var input models.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := h.service.CreateItem(&input); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c, http.StatusCreated, "Item created successfully", input)
}
