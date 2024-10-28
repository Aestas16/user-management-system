package main

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo"

    "user-management-system/internal/config"
    "user-management-system/internal/model"
)

func main() {
    config.Init();
    model.InitDB();
}