package main

import (
	"context"
	"log"
	"os"

	"dummies-backend/internal/api/handler"
	"dummies-backend/internal/api/router"
	"dummies-backend/internal/application/usecase"
	"dummies-backend/internal/infrastructure/persistence"
)

func main() {
	ctx := context.Background()

	pool, err := persistence.NewDB(ctx)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	userRepo := persistence.NewUserRepository(pool)
	projectRepo := persistence.NewProjectRepository(pool)
	dummyDataRepo := persistence.NewDummyDataRepository(pool)

	userUC := usecase.NewUserUseCase(userRepo)
	projectUC := usecase.NewProjectUseCase(projectRepo)
	dummyDataUC := usecase.NewDummyDataUseCase(dummyDataRepo, projectRepo)

	healthHandler := handler.NewHealthHandler()
	userHandler := handler.NewUserHandler(userUC)
	projectHandler := handler.NewProjectHandler(projectUC)
	dummyDataHandler := handler.NewDummyDataHandler(dummyDataUC)

	r := router.NewRouter(healthHandler, userHandler, projectHandler, dummyDataHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
