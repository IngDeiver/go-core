package files

import (
	middlewares "github.com/ingdeiver/go-core/src/auth/infrastructure/framework/middlewares"
	"github.com/ingdeiver/go-core/src/config"
	services "github.com/ingdeiver/go-core/src/files/application/services"
	domain "github.com/ingdeiver/go-core/src/files/domain/interfaces"
	controllers "github.com/ingdeiver/go-core/src/files/infrastructure/controllers"
	repositories "github.com/ingdeiver/go-core/src/files/infrastructure/s3/repositories"
)


var FilesRepository domain.BaseFileRepository
var FilesService domain.BaseFileService
var FilesController *controllers.FilesController

// Instance Repositories, Services, Controllers and more
func InitFilesModule(){
	router := config.GetRouter()

	// ----- repositories -----
	FilesRepository = repositories.New()


	// ----- services -----
	FilesService = services.New(FilesRepository)

	// ----- controllers -----
	FilesController = controllers.New(FilesService)

	userRouter := router.Group("/files")
	{
		userRouter.Use(middlewares.AuthMiddleware())
		userRouter.POST("", FilesController.Upload)
		userRouter.GET(":key", FilesController.Get)
		userRouter.DELETE(":key", FilesController.Remove)
	}
}