package handler

import (
	"github.com/khdiyz/book-store/model"
	handlerfunc "github.com/khdiyz/book-store/utils/handler_func"
	"github.com/khdiyz/book-store/utils/response"
	"github.com/labstack/echo/v4"
)

const (
	searchKey = "search"
	id        = "id"
)

// CreateCategory
// @Description Create Category
// @Summary Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Param create body model.CreateCategory true "Create Category"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/categories [post]
func (h *Handler) createCategory(c echo.Context) error {
	var body model.CreateCategory

	if err := c.Bind(&body); err != nil {
		return response.HandleResponse(c, response.BadRequest, err, nil, nil)
	}

	id, err := h.services.CategoryI.Create(body)
	if err != nil {
		return response.HandleResponse(c, response.InternalServerError, err, nil, nil)
	}

	return response.HandleResponse(c, response.Created, nil, map[string]interface{}{
		"id": id,
	}, nil)
}

// GetCategoryList
// @Description Get Category List
// @Summary Get Category List
// @Tags Category
// @Accept json
// @Produce json
// @Param search  query string false "searchKey"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/categories [get]
func (h *Handler) getCategoryList(c echo.Context) error {
	var request model.CategoryListRequest

	request.SearchKey = handlerfunc.GetNullStringQuery(c, searchKey)

	categories, err := h.services.CategoryI.GetList(request)
	if err != nil {
		return response.HandleResponse(c, response.InternalServerError, err, nil, nil)
	}

	return response.HandleResponse(c, response.OK, nil, categories, nil)
}

// GetCategory
// @Description Get Category
// @Summary Get Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id  path string false "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/categories/{id} [get]
func (h *Handler) getCategoryByID(c echo.Context) error {
	id, err := handlerfunc.GetIntParam(c, id)
	if err != nil {
		response.HandleResponse(c, response.BadRequest, err, nil, nil)
	}

	category, err := h.services.CategoryI.GetByID(model.GetCategoryRequest{ID: id})
	if err != nil {
		response.HandleResponse(c, response.Internal, err, nil, nil)
	}

	return response.HandleResponse(c, response.OK, nil, category, nil)
}

func (h *Handler) updateCategory(c echo.Context) error {

	return nil
}

func (h *Handler) deleteCategory(c echo.Context) error {

	return nil
}
