package handler

import (
	"fmt"
	"go-echo/entity"
	"go-echo/services"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type biodataHandler struct {
	biodataService services.Service
}

func NewBiodataHandler(biodataService services.Service) *biodataHandler {
	return &biodataHandler{biodataService}
}

func Root(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func (h *biodataHandler) GetAll(c echo.Context) error {
	biodatas, err := h.biodataService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	var biodatasResponse []entity.Response

	for _, b := range biodatas {
		biodataResponse := convertToBiodataResponse(b)
		biodatasResponse = append(biodatasResponse, biodataResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": biodatasResponse,
	})
}

func (h *biodataHandler) FindByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}

	biodata, err := h.biodataService.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if biodata == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "Biodata not found",
		})
	}

	biodataResponse := convertToBiodataResponse(biodata)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": biodataResponse,
	})
}

func (h *biodataHandler) Create(c echo.Context) error {
	biodataRequest := new(entity.Request)
	err := c.Bind(biodataRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
	}

	err = h.biodataService.Create(biodataRequest)
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			report := struct {
				Message string `json:"message"`
			}{}
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s is required",
					err.Field())
			case "gte":
				report.Message = fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param())
			case "lte":
				report.Message = fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param())
			default:
				report.Message = "Invalid request payload"
			}

			return c.JSON(http.StatusBadRequest, report)
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Biodata created successfully",
	})
}

func (h *biodataHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}

	var biodataUpdate entity.Update
	if err := c.Bind(&biodataUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
	}

	biodata, err := h.biodataService.Update(id, &biodataUpdate)
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			report := struct {
				Message string `json:"message"`
			}{}
			switch err.Tag() {
			case "gte":
				report.Message = fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param())
			case "lte":
				report.Message = fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param())
			default:
				report.Message = "Invalid request payload"
			}

			return c.JSON(http.StatusBadRequest, report)
		}
	}

	if biodata == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "Biodata not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Biodata updated successfully",
		"data":    convertToBiodataResponse(biodata),
	})
}

func (h *biodataHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}

	err = h.biodataService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Biodata deleted succesfully",
	})
}

func convertToBiodataResponse(b *entity.Biodata) entity.Response {
	return entity.Response{
		ID:        b.ID,
		NAME:      b.NAME,
		AGE:       b.AGE,
		ADDRESS:   b.ADDRESS,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}
