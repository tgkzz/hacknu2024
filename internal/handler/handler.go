package handler

import (
	"backend/internal/service"
	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) GenerateRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.Static("/static/swagger", "./docs")

	r.GET("/helloworld", h.HelloWorld)

	v1 := r.Group("/v1")
	{
		swagger := v1.Group("/swagger")
		{
			url := ginSwagger.URL("http://localhost:9090/static/swagger/swagger.json")
			swagger.GET("*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
		}
		student := v1.Group("/student")
		{
			student.GET("/get-student-id-by-name", h.GetUserByName)
		}
		bot := v1.Group("/bot")
		{
			bot.POST("/answer-student", h.RequestStudentQuestion)
		}
		db := v1.Group("/req")
		{
			db.Group("/get-info")
		}

	}

	return r
}
