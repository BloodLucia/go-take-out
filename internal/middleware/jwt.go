package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	jwtPkg "github.com/golang-jwt/jwt/v5"
	myErrs "github.com/kalougata/go-take-out/pkg/errors"
	"github.com/kalougata/go-take-out/pkg/jwt"
	"github.com/kalougata/go-take-out/pkg/response"
)

type JWTMiddleware struct {
	jwt *jwt.JWT
}

func NewJWTMiddleware(jwt *jwt.JWT) *JWTMiddleware {
	return &JWTMiddleware{jwt}
}

func (jm *JWTMiddleware) JWTAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return response.Build(c, myErrs.ErrUnauthorized(), nil)
		}
		claims, err := jm.jwt.ParseToken(tokenString)
		if err != nil {
			if errors.Is(err, jwtPkg.ErrTokenExpired) {
				return response.Build(c, myErrs.ErrUnauthorized().WithMsg("token已过期"), nil)
			}
			return response.Build(c, myErrs.ErrUnauthorized().WithMsg("token校验失败"), nil)
		}

		c.Set("userId", claims.UserId)
		return c.Next()
	}
}
