package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Venukishore-R/CODECRAFT_BW_03/config"
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	_ "github.com/joho/godotenv/autoload"
)

type Claims struct {
	Id    uint
	Email string
	Role  string
	jwt.StandardClaims
}

func NewClaims(id uint, email string, role string, sc jwt.StandardClaims) *Claims {
	return &Claims{
		Id:    id,
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  sc.IssuedAt,
			ExpiresAt: sc.ExpiresAt,
		},
	}
}

type Auth struct {
	Claims *Claims
	Key    string
}

func NewAuth(user *models.User, sC jwt.StandardClaims, key string) *Auth {
	return &Auth{
		Claims: NewClaims(user.Id, user.Email, user.Role, sC),
		Key:    key,
	}
}

type AuthInterface interface {
	GenerateToken() (string, error)
}

func (a *Auth) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.Claims)
	return token.SignedString([]byte(a.Key))
}

func GeneralAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getTokenFromHeaders(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error(), "message": "unauthorized user"})
			return
		}

		userClaims, err := parseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   err.Error(),
				"message": "unauthorized user, error parsing token",
			})

			return
		}

		c.Set("user", userClaims)

		c.Next()
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getTokenFromHeaders(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error(), "message": "unauthorized user"})
			return
		}

		claims, err := parseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   err.Error(),
				"message": "unauthorized user, error parsing token",
			})

			return
		}
		if claims.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "only admin is allowed to access this endpoints",
				"message": "unauthorized user, insufficient permissions",
			})
			return
		}

		c.Set("user", claims)

		c.Next()
	}
}
func getTokenFromHeaders(c *gin.Context) (string, error) {
	bearerToken := c.Request.Header.Get("Authorization")

	if !strings.HasPrefix(bearerToken, "Bearer ") {
		return "", fmt.Errorf("missing bearer token")
	}

	token := strings.TrimPrefix(bearerToken, "Bearer ")
	return token, nil
}

func parseToken(token string) (*Claims, error) {
	config, _ := config.LoadConfig()

	tok, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Key), nil
	})

	return tok.Claims.(*Claims), err
}
