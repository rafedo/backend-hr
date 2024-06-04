package Middleware

import (
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"muhammadiyah/Constant"
	"muhammadiyah/Model/Domain"
	"muhammadiyah/Model/Web"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func JwtMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var acces_token string
		// Ambil token JWT dari header Authorization
		authHeader := c.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			acces_token = strings.TrimPrefix(authHeader, "Bearer ")
		}
		tokenClaim, err := ValidateTokenJwt(acces_token, os.Getenv("ACCESS_TOKEN_PUBLIC_KEY"))
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(Web.ErrorResponse(Constant.AuthorizeError, nil))

		}

		// Set data dari token ke dalam lokal untuk digunakan dalam handler selanjutnya
		c.Locals("userID", tokenClaim.UserID)
		return c.Next()
	}
}

func CheckPermissions(allowedPosition []string, allowedDepartment []string, allowedPlacement []string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		userData := c.Locals("token")
		fmt.Println(userData)
		if userData == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized access",
			})
		}

		user, ok := userData.(*Domain.User)
		if ok == false {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized access",
			})
		}
		if user == nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Anda tidak memiliki akses ke wilayah ini",
			})
		}
		// Memeriksa penempatan, departemen, dan jabatan
		if !checkPlacement(user.Penempatan, allowedPlacement) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Anda tidak memiliki akses ke wilayah ini",
			})
		}

		if !checkDepartment(user.Departemen, allowedDepartment) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Anda tidak memiliki akses ke departemen ini",
			})
		}

		if !checkPosition(user.Jabatan, allowedPosition) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Anda tidak memiliki jabatan yang sesuai",
			})
		}

		// Lanjutkan eksekusi jika pengguna memiliki izin yang diperlukan
		return c.Next()
	}
}

func checkPlacement(userPlacement string, allowedPlacements []string) bool {
	// Memeriksa apakah pengguna memiliki akses ke penempatan yang diperlukan
	for _, allowed := range allowedPlacements {
		if allowed == "all" {
			return true // Mengizinkan semua departemen
		}
		if userPlacement == allowed {
			return true
		}
	}
	return false
}

func checkDepartment(userDepartment string, allowedDepartments []string) bool {
	// Memeriksa apakah pengguna memiliki akses ke departemen yang diperlukan
	for _, allowed := range allowedDepartments {
		if allowed == "all" {
			return true // Mengizinkan semua departemen
		}
		if userDepartment == allowed {
			return true
		}
	}
	return false
}

func checkPosition(userPosition string, allowedPositions []string) bool {
	// Memeriksa apakah pengguna memiliki jabatan yang diperlukan
	for _, allowed := range allowedPositions {
		if allowed == "all" {
			return true // Mengizinkan semua departemen
		}
		if userPosition == allowed {
			return true
		}
	}
	return false
}

func GenerateTokenJwtV2(jwTokenTime time.Duration, user int64, privateKey string) (*Domain.JwtTokenDetail, error) {
	expTime := time.Now()

	tokenDetail := &Domain.JwtTokenDetail{
		ExpiresIn: new(int64),
		Token:     new(string),
	}

	*tokenDetail.ExpiresIn = expTime.Add(jwTokenTime).Unix()
	tokenDetail.UserID = user

	decodePrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodePrivateKey)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	claimJwt := jwt.MapClaims{}
	claimJwt["sub"] = user
	claimJwt["exp"] = tokenDetail.ExpiresIn
	claimJwt["iat"] = expTime.Unix()
	claimJwt["nbf"] = expTime.Unix()
	*tokenDetail.Token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, claimJwt).SignedString(key)
	if err != nil {
		return nil, err
	}

	return tokenDetail, nil

}

func ValidateTokenJwt(token string, publicKey string) (*Domain.JwtTokenDetail, error) {
	decodePublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, err
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(decodePublicKey)
	if err != nil {
		return nil, err
	}
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok && !parsedToken.Valid {
		return nil, err
	}

	subStr := fmt.Sprint(claims["sub"])
	userID, err := strconv.ParseInt(subStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid userID: %v", err)
	}
	return &Domain.JwtTokenDetail{
		UserID: userID,
	}, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // Ganti "secret" dengan kunci rahasia yang digunakan untuk menandatangani token
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}
