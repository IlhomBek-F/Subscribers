package route

import (
	"net/http"
	"subscribers/api/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger" // add this import
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

type Server struct {
	Port int
	DB   *gorm.DB
}

func (s Server) RegisterRoute() http.Handler {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(10))))
	e.Use(configureCORS())

	server := controller.Server{
		DB: s.DB,
	}

	subscriberGroup := e.Group("/api/subscriber")

	subscriberGroup.GET("/list", server.GetUsers)
	subscriberGroup.POST("/create", server.CreateNewUser)
	subscriberGroup.PUT("/update", server.UpdateUser)
	subscriberGroup.DELETE("/delete/:id", server.DeleteUser)
	subscriberGroup.GET("/:id", server.GetUserById)
	subscriberGroup.GET("/cost", server.CalculateSubsCost)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}

func configureCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           30,
	})
}
