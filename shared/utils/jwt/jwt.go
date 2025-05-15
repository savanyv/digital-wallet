package jwt

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/savanyv/digital-wallet/shared/config"
)

type JWTClaim struct {
	UserID string `json:"user_id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateToken(userID uint, email string) (string, error)
	ValidateToken(tokenString string) (*JWTClaim, error)
}

type jwtService struct {
	secretKey string
	issuer string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: config.LoadConfig().SecretKey,
		issuer: "digital-wallet",
	}
}

func (j *jwtService) GenerateToken(userID int64, email string) (string, error) {
	claims := &JWTClaim{
		UserID: strconv.Itoa(int(userID)),
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: j.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(tokenString string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	return claims, nil
}

func JWTMiddleware(jwtService JWTService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing token",
			})
		}

		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Simpan klaim di context
		c.Locals("user", claims)
		return c.Next()
	}
}
