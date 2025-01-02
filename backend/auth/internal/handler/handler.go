package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/rajivgeraev/flippy-toys/backend/auth/internal/service"
)

type Handler struct {
    services *service.Services
}

func NewHandler(services *service.Services) *Handler {
    return &Handler{
        services: services,
    }
}

func (h *Handler) InitRoutes() *gin.Engine {
    router := gin.Default()

    api := router.Group("/api/v1")
    {
        users := api.Group("/users")
        {
            users.POST("", h.createUser)
            users.GET("/:id", h.getUser)
            users.PUT("/:id", h.updateUser)
        }
    }

    return router
}