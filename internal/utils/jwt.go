package utils

import (
    "time"
    "errors"
    "github.com/golang-jwt/jwt/v4"
    "github.com/labstack/echo"

    "user-management-system/internal/model"
    "user-management-system/internal/config"
)

var jwtKey = []byte(config.Config.Server.JwtKey)

type Claims struct {
    user model.User
    isAdmin bool
    jwt.RegisteredClaims
}

func generateToken(user model.User, isAdmin bool) (string, error) {
    expirationTime := time.Now().Add(5 * time.Minute)
    claims := &Claims{
        user: user,
        isAdmin: isAdmin,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    return tokenString, err
}

func parseToken(tokenString string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if !token.Valid || err != nil {
        return claims, errors.New("invalid token")
    }
    if time.Until(claims.ExpiresAt.Time) < 0 {
        return claims, errors.New("token expired")
    }
    return claims, err
}

func jwtAuthMiddleware() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            tokenString := c.Request().Header.Get("Authorization")
            if tokenString == "" {
                return echo.NewHTTPError(401, "token not found")
            }
            claims, err := parseToken(tokenString)
            if err == errors.New("invalid token") {
                return echo.NewHTTPError(401, "invalid token")
            } else if err == errors.New("token expired") {
                return echo.NewHTTPError(401, "token expired")
            } else if err != nil {
                return echo.ErrInternalServerError
            }
            c.Set("claims", claims)
            return next(c);
        }
    }
}