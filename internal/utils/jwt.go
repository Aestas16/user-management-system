package utils

import (
    "errors"
    "github.com/golang-jwt/jwt/v4"

    "user-management-system/internal/model"
    "user-management-system/internal/config"
)

var jwtKey = []byte(config.Config.JwtKey)

type Claims struct {
    user model.User
    jwt.RegisteredClaims
}

func generateToken(user model.User) (string, err) {
    expirationTime := time.Now().Add(5 * time.Minute)
    claims := &Claims{
        user: user,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime)
        }
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    return tokenString, err
}

func parseToken(tokenString string) (Claims, err) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if !token.Valid || err != nil {
        return nil, errors("invalid token")
    }
    return claims, err
}