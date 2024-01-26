package handlers

import (
	"testAPI/pkg/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service service.Service
}

func NewHandler(service *service.Service) *handler {
	return &handler{
		service: *service,
	}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/", h.GetPeople)
		api.POST("/", h.CreatePerson)
		api.DELETE("/:id", h.DeletePerson)
		api.PUT("/:id", h.UpdatePerson)
	}

	return router
}
