package main

import (
    "fmt"
    "github.com/labstack/echo"

    "user-management-system/internal/config"
    "user-management-system/internal/model"
    "user-management-system/internal/router"
)

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    config.InitConfig()
    model.InitDB()
    router.InitRouter(e)
    e.Logger.Fatal(e.Start(":" + config.Config.Server.Port))
}