package handler

import (
	services "github.com/Ki4EH/lib-service/account/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func New(srv *services.Service) *Handler {
	return &Handler{services: srv}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		api.POST("register", h.register)
		api.POST("login", h.login)
		api.POST("logout", h.logout)
		api.GET("verify/:token", h.verify)
	}

	return router
}
