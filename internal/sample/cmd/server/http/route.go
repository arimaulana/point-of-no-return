package http

import (
	"github.com/arimaulana/point-of-no-return/internal/common/pkg/log"
	"github.com/arimaulana/point-of-no-return/internal/sample/internal/healthcheck"
	"github.com/arimaulana/point-of-no-return/internal/sample/internal/user"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// Routing setup api routing
func Routing(app *fiber.App, db *sqlx.DB, logger log.Logger) {
	validate = validator.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "success",
			"error":   false,
		})
	})

	// register no auth api path group
	healthcheck.RegisterHandlers(app)

	// init services
	userRepo := user.NewMysqlRepo(db, logger)
	userSvc := user.NewService(userRepo, logger)
	user.RegisterNoAuthHttpHandlers(app, user.NewNoAuthHttpHandler(userSvc, logger))

	// register api path group
	apiV1 := app.Group("/api/v1")

	userApiV1 := apiV1.Group("/users")
	user.RegisterHttpHandlers(userApiV1, user.NewApiHttpHandler(userSvc, logger))

	// register admin api path group
	adminV1 := app.Group("/admin/v1")

	userAdminV1 := adminV1.Group("/users")
	user.RegisterAdminHttpHandlers(userAdminV1, user.NewAdminHttpHandler(userSvc, logger))

}
