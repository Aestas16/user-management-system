package controller

import (
    "github.com/labstack/echo/v4"

    "user-management-system/internal/config"
    "user-management-system/internal/model"
)

func userInfo(c echo.Context) error {
    id := c.Param("id")
    user, err := model.findUserById(id)
    if err == model.userNotFound {
        return echo.ErrNotFound
    } else if err != nil {
        return echo.ErrInternalServerError
    }
    var resp struct {
        username    string
        email       string
    }
    resp.username = user.username
    resp.email = user.email
    return c.JSON(200, &resp);
}