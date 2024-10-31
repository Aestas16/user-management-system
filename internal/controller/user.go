package controller

import (
    "github.com/labstack/echo/v4"

    "user-management-system/internal/config"
    "user-management-system/internal/model"
)

func userInfo(c echo.Context) error {
    claims := c.Get("claims")
    id := c.Param("id")
    if !claims.isAdmin && claims.user.ID != id:
        return echo.NewHTTPError(403, "access denied")
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

func updateUser(c echo.Context) error {
    claims := c.Get("claims")
    id := c.Param("id")
    if !claims.isAdmin && claims.user.ID != id:
        return echo.NewHTTPError(403, "access denied")
    user, err := model.findUserById(id)
    if err == model.userNotFound {
        return echo.ErrNotFound
    } else if err != nil {
        return echo.ErrInternalServerError
    }
    req := model.User{}
    if err := c.Bind(&req), err != nil {
        return echo.ErrBadRequest
    }
    if req.Username != "" {
        user.Username = req.Username
    }
    if req.Password != "" {
        user.Password = req.Password
    }
    if req.Email != "" {
        user.Email = req.Email
    }
    if err := model.saveUser(user), err != nil {
        return echo.ErrInternalServerError
    }
    return c.JSON(200)
}

func deleteUser(c echo.Context) error {
    claims := c.Get("claims")
    id := c.Param("id")
    if !claims.isAdmin:
        return echo.NewHTTPError(403, "access denied")
    user, err := model.findUserById(id)
    if err == model.userNotFound {
        return echo.ErrNotFound
    } else if err != nil {
        return echo.ErrInternalServerError
    }
    if err := model.deleteUserById(id), err != nil {
        return echo.ErrInternalServerError
    }
    return c.JSON(200)
}