package router

import (
	"dummies-backend/internal/api/handler"
	"dummies-backend/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	healthHandler *handler.HealthHandler,
	userHandler *handler.UserHandler,
	projectHandler *handler.ProjectHandler,
	dummyDataHandler *handler.DummyDataHandler,
) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	r.GET("/health", healthHandler.Health)

	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware())
	{
		v1.GET("/user", userHandler.GetProfile)
		v1.PUT("/user", userHandler.UpdateProfile)
		v1.DELETE("/user", userHandler.DeleteAccount)

		v1.GET("/projects", projectHandler.List)
		v1.POST("/projects", projectHandler.Create)
		v1.GET("/projects/:id", projectHandler.Get)
		v1.PUT("/projects/:id", projectHandler.Update)
		v1.DELETE("/projects/:id", projectHandler.Delete)

		v1.GET("/projects/:id/dummies", dummyDataHandler.List)
		v1.POST("/projects/:id/dummies", dummyDataHandler.Create)
		v1.GET("/dummies/:uuid", dummyDataHandler.Get)
		v1.PUT("/dummies/:uuid", dummyDataHandler.Update)
		v1.DELETE("/dummies/:uuid", dummyDataHandler.Delete)
	}

	return r
}
