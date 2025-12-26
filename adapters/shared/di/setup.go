package di

import (
	"fmt"
	"hinsun-backend/adapters/secondary/repositories"
	"hinsun-backend/internal/domain/applications"
	"hinsun-backend/internal/domain/services"

	"gorm.io/gorm"
)

func SetupDependencies(db *gorm.DB) {
	// Initialize and wire up dependencies here
	experienceRepository := repositories.NewExperienceRepostory(db)

	// Initialize domain services
	experienceService := services.NewExperienceService(experienceRepository)

	// Initialize application services
	globalAppService := applications.NewGlobalAppService(experienceService)
	fmt.Println(globalAppService)
}
