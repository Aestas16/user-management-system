package router

import (
    "fmt"
    "github.com/labstack/echo/v4"

    "user-management-system/internal/config"
    "user-management-system/internal/controller"
    "user-management-system/internal/utils"
)

func InitRouter(e *echo.Echo) {
    apiVersion := config.Config.Server.Version
    apiUser := e.Group(fmt.Sprintf("/%s/user", apiVersion))
    apiUser.Use(utils.JWTAuthMiddleware())
    apiUser.GET("/:id", controller.userInfo)
    apiUser.POST("/:id", controller.userInfo)
    apiUser.POST("/:id/update", controller.updateUser)
    apiUser.POST("/:id/delete", controller.deleteUser)
    apiUser.POST("/token", controller.refreshToken)
    api := e.Group(fmt.Sprintf("/%s", apiVersion))
    api.POST("/login", controller.loginUser)
    api.POST("/register", controller.registerUser)
}