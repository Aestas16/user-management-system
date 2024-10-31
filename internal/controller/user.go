package controller

import (
    "fmt"
    "time"
    "crypto/md5"
    "github.com/labstack/echo"

    "user-management-system/internal/config"
    "user-management-system/internal/model"
    "user-management-system/internal/utils"
)

type User struct {
    Username    string
    Password    string
    Email       string
}

func userInfo(c echo.Context) error {
    claims := c.Get("claims")
    id := c.Param("id")
    if !claims.isAdmin && claims.user.ID != id {
        return echo.NewHTTPError(403, "access denied")
    }
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
    return c.JSON(200, &resp)
}

func updateUser(c echo.Context) error {
    claims := c.Get("claims")
    id := c.Param("id")
    if !claims.isAdmin && claims.user.ID != id {
        return echo.NewHTTPError(403, "access denied")
    }
    user, err := model.findUserById(id)
    if err == model.userNotFound {
        return echo.ErrNotFound
    } else if err != nil {
        return echo.ErrInternalServerError
    }
    req := User{}
    if err := c.Bind(&req), err != nil {
        return echo.ErrBadRequest
    }
    user.Username = req.Username
    user.Password = fmt.Sprintf("%x", md5.Sum([]byte(req.Password)))
    user.Email = req.Email
    if err := model.saveUser(user), err != nil {
        return echo.ErrInternalServerError
    }
    var resp struct {
        message string
    }
    resp.message = "Success!"
    return c.JSON(200, &resp)
}

func deleteUser(c echo.Context) error {
    claims := c.Get("claims")
    id := c.Param("id")
    if !claims.isAdmin {
        return echo.NewHTTPError(403, "access denied")
    }
    user, err := model.findUserById(id)
    if err == model.userNotFound {
        return echo.ErrNotFound
    } else if err != nil {
        return echo.ErrInternalServerError
    }
    if err := model.deleteUserById(id), err != nil {
        return echo.ErrInternalServerError
    }
    var resp struct {
        message string
    }
    resp.message = "Success!"
    return c.JSON(200, &resp)
}

func registerUser(c echo.Context) error {
    req := User{}
    if err := c.Bind(&req), err != nil {
        return echo.ErrBadRequest
    }
    if req.Username == "" || req.Password == "" {
        return echo.ErrBadRequest
    }
    user := model.User{}
    user.Username = req.Username
    user.Password = fmt.Sprintf("%x", md5.Sum([]byte(req.Password)))
    user.Email = req.Email
    err := model.createUser(user)
    if err == model.userAlreadyExist {
        return echo.NewHTTPError(403, "access denied")
    } else if err != nil {
        return echo.ErrInternalServerError
    }
    var resp struct {
        message string
    }
    resp.message = "Success!"
    return c.JSON(201, &resp)
}

func loginUser(c echo.Context) error {
    var req struct {
        Username    string
        Password    string
    }
    if err := c.Bind(&req), err != nil {
        return echo.ErrBadRequest
    }
    if req.Username == config.Config.Server.Admin.Username && req.Password == config.Config.Server.Admin.Password {
        tokenString, err := utils.generateToken(model.User{}, true)
        if err != nil {
            return echo.ErrInternalServerError
        }
        var resp struct {
            tokenString string
        }
        resp.tokenString = tokenString
        return c.JSON(200, &resp)
    }
    req.Password = fmt.Sprintf("%x", md5.Sum([]byte(req.Password)))
    user, err := model.findUserByName(req.Username)
    if err == model.userNotFound {
        return echo.NewHTTPError(403, "user not found")
    } else if err != nil {
        return echo.ErrInternalServerError
    }
    if user.Password != req.Password {
        return echo.NewHTTPError(403, "wrong password")
    }
    tokenString, err := utils.generateToken(user, false)
    if err != nil {
        return echo.ErrInternalServerError
    }
    var resp struct {
        tokenString string
    }
    resp.tokenString = tokenString
    return c.JSON(200, &resp)
}

func refreshToken(c echo.Context) error {
    claims := c.Get("claims")
    expirationTime := time.Now().Add(5 * time.Minute)
    claims.ExpiresAt.Time = expirationTime
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(utils.jwtKey)
    if err != nil {
        return echo.ErrInternalServerError
    }
    var resp struct {
        tokenString string
    }
    resp.tokenString = tokenString
    return c.JSON(200, &resp)
}