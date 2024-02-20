package config

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/krostyle/auth-systme-go/src/adapters/controller"
	"github.com/krostyle/auth-systme-go/src/application/usecase"

	"github.com/krostyle/auth-systme-go/src/infrastructure/api/router"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/configuration"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/persistence"
	"github.com/krostyle/auth-systme-go/src/infrastructure/service"
)

func Setup(app *fiber.App) {

	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			panic(fmt.Sprintf("Error loading .env file: %v", err))
		}
	}

	gormDB, err := configuration.NewPostgresDB()
	if err != nil {
		panic(fmt.Sprintf("Error connecting to database: %v", err))
	}

	identifierGenerator := service.NewUUIDGenerator()
	passwordHasher := service.NewPasswordHasher()

	healthCheckController := controller.NewHealthCheckController()

	// permissionRepository := persistence.NewPermissionRepository(gormDB)
	// permissionCrud := crud.NewermissionCrud(permissionRepository)
	// permissionController := controller.NewPermissionController(permissionCrud)

	// roleRepository := persistence.NewRoleRepository(gormDB)
	// roleCrud := crud.NewRoleCrud(roleRepository)
	// roleController := controller.NewRoleController(roleCrud)

	userRepository := persistence.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository, identifierGenerator, passwordHasher)
	userController := controller.NewUserController(userUseCase)

	fmt.Println("Setting up routes")
	router.SetupRouter(
		app,
		// permissionController,
		// roleController,
		userController,
		healthCheckController)

}
