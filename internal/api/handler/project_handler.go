package handler

import (
	"dummies-backend/internal/api/dto"
	"dummies-backend/internal/application/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	uc *usecase.ProjectUseCase
}

func NewProjectHandler(uc *usecase.ProjectUseCase) *ProjectHandler {
	return &ProjectHandler{uc: uc}
}

func (h *ProjectHandler) List(c *gin.Context) {
	userID := c.GetString("userID")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	projects, totalCount, err := h.uc.ListByUser(c.Request.Context(), userID, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var items []*dto.ProjectResponse
	for _, p := range projects {
		items = append(items, dto.ToProjectResponse(p))
	}
	if items == nil {
		items = []*dto.ProjectResponse{}
	}

	c.JSON(http.StatusOK, dto.ProjectListResponse{
		Projects:   items,
		TotalCount: totalCount,
		Page:       page,
		PerPage:    usecase.ProjectsPerPage,
	})
}

func (h *ProjectHandler) Get(c *gin.Context) {
	userID := c.GetString("userID")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
		return
	}

	project, err := h.uc.GetByID(c.Request.Context(), id, userID)
	if err != nil {
		status := http.StatusNotFound
		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToProjectResponse(project))
}

func (h *ProjectHandler) Create(c *gin.Context) {
	userID := c.GetString("userID")

	var req dto.CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := h.uc.Create(c.Request.Context(), userID, req.Name, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToProjectResponse(project))
}

func (h *ProjectHandler) Update(c *gin.Context) {
	userID := c.GetString("userID")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
		return
	}

	var req dto.UpdateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.Update(c.Request.Context(), id, userID, req.Name, req.Description); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		} else if err.Error() == "project not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (h *ProjectHandler) Delete(c *gin.Context) {
	userID := c.GetString("userID")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
		return
	}

	if err := h.uc.Delete(c.Request.Context(), id, userID); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		} else if err.Error() == "project not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
