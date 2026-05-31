package handler

import (
	"dummies-backend/internal/api/dto"
	"dummies-backend/internal/application/usecase"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DummyDataHandler struct {
	uc *usecase.DummyDataUseCase
}

func NewDummyDataHandler(uc *usecase.DummyDataUseCase) *DummyDataHandler {
	return &DummyDataHandler{uc: uc}
}

func (h *DummyDataHandler) List(c *gin.Context) {
	userID := c.GetString("userID")
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
		return
	}

	list, err := h.uc.ListByProject(c.Request.Context(), projectID, userID)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		} else if err.Error() == "project not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	var items []*dto.DummyDataResponse
	for _, d := range list {
		items = append(items, dto.ToDummyDataResponse(d))
	}
	if items == nil {
		items = []*dto.DummyDataResponse{}
	}

	c.JSON(http.StatusOK, items)
}

func marshalStringSlice(s []string) string {
	b, _ := json.Marshal(s)
	return string(b)
}

func (h *DummyDataHandler) Create(c *gin.Context) {
	userID := c.GetString("userID")
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
		return
	}

	var req dto.CreateDummyDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	colNameJSON := marshalStringSlice(req.ColumnName)
	colTypeJSON := marshalStringSlice(req.ColumnType)
	var colValidateJSON *string
	if len(req.ColumnValidate) > 0 {
		v := marshalStringSlice(req.ColumnValidate)
		colValidateJSON = &v
	}

	data, err := h.uc.Create(c.Request.Context(), projectID, userID, req.TableName, colNameJSON, colTypeJSON, colValidateJSON)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		} else if err.Error() == "project not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToDummyDataResponse(data))
}

func (h *DummyDataHandler) Get(c *gin.Context) {
	userID := c.GetString("userID")
	uuid := c.Param("uuid")

	data, err := h.uc.GetByID(c.Request.Context(), uuid, userID)
	if err != nil {
		status := http.StatusNotFound
		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToDummyDataResponse(data))
}

func (h *DummyDataHandler) Update(c *gin.Context) {
	userID := c.GetString("userID")
	uuid := c.Param("uuid")

	var req dto.UpdateDummyDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	colNameJSON := marshalStringSlice(req.ColumnName)
	colTypeJSON := marshalStringSlice(req.ColumnType)
	var colValidateJSON *string
	if len(req.ColumnValidate) > 0 {
		v := marshalStringSlice(req.ColumnValidate)
		colValidateJSON = &v
	}

	if err := h.uc.Update(c.Request.Context(), uuid, userID, req.TableName, colNameJSON, colTypeJSON, colValidateJSON); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		} else if err.Error() == "dummy data not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (h *DummyDataHandler) Delete(c *gin.Context) {
	userID := c.GetString("userID")
	uuid := c.Param("uuid")

	if err := h.uc.Delete(c.Request.Context(), uuid, userID); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		} else if err.Error() == "dummy data not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
