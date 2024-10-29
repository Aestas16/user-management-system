package router

import (
    "fmt"
    "github.com/labstack/echo/v4"

    "user-management-system/internal/config"
    "user-management-system/internal/controller"
)

func InitRouter() {
    apiVersion := config.Config.Server.Version
    e.GET(fmt.Sprintf("/%s/user/:id", apiVersion), controller.userInfo)
    e.POST(fmt.Sprintf("/%s/user/:id", apiVersion), controller.userInfo)
}