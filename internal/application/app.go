package application

import (
	"fmt"
	"log"

	"github.com/MostajeranMohammad/blog/config"
	_ "github.com/MostajeranMohammad/blog/docs"
	"github.com/MostajeranMohammad/blog/internal/controller"
	"github.com/MostajeranMohammad/blog/internal/entity"
	"github.com/MostajeranMohammad/blog/internal/repo"
	"github.com/MostajeranMohammad/blog/internal/routes"
	"github.com/MostajeranMohammad/blog/internal/usecase"
	"github.com/MostajeranMohammad/blog/pkg/guards"
	"github.com/MostajeranMohammad/blog/pkg/logger"
	"github.com/MostajeranMohammad/blog/pkg/utils"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func Run(cfg *config.Config) {
	l, zL := logger.New(cfg.Log.Level)

	// Repository
	dsn := cfg.PG.DSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("failed to connect database")
	}

	sqlDB, err := db.DB()
	defer func() {
		sqlDB.Close()
		if err != nil {
			log.Fatalln("failed to close sqlDB")
		}
	}()

	if err != nil {
		log.Fatalln("failed to extract sqlDB")
	}

	// auto migrate models
	db.AutoMigrate(&entity.BlogPost{})

	// initialize repos
	blogPostRepo := repo.NewBlogPostRepo(db)

	// initialize usecases
	blogPostUsecase := usecase.NewBlogPostUsecase(blogPostRepo)

	// initialize controllers
	blogPostController := controller.NewBlogPostController(blogPostUsecase)

	// initialize guards
	jwtGuard := guards.NewJWTGuard(cfg.JwtSecret)

	// HTTP Server
	httpApp := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: utils.FiberErrorHandler(l),
	})
	httpApp.Use(fiberzerolog.New(fiberzerolog.Config{Logger: zL}))
	httpApp.Get("/swagger/*", swagger.HandlerDefault)
	httpApp.Mount("/blog-post", routes.NewStaticFileRouter(blogPostController, jwtGuard))

	httpApp.Listen(fmt.Sprintf(":%s", cfg.HTTP.Port))
}
