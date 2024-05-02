package Middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"muhammadiyah/Model/Domain"
	"net/http"
	"strings"
	"time"
)

func JwtMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil token JWT dari header Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing Authorization header",
			})
		}

		// Periksa apakah header Authorization memiliki format yang sesuai
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Authorization header format",
			})
		}

		// Ambil token JWT dari bagian kedua header Authorization
		tokenString := tokenParts[1]

		// Verifikasi token JWT
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil // Ganti "secret" dengan kunci rahasia yang digunakan untuk menandatangani token
		})
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Failed to parse JWT token",
			})
		}

		if !token.Valid {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid JWT token",
			})
		}

		// Ambil data dari klaim token JWT
		userID := claims["user_id"].(string)
		jabatan := claims["jabatan"].(string)
		departemen := claims["departemen"].(string)
		penempatan := claims["penempatan"].(string)

		// Set data dari token ke dalam lokal untuk digunakan dalam handler selanjutnya
		c.Locals("userID", userID)
		c.Locals("jabatan", jabatan)
		c.Locals("departemen", departemen)
		c.Locals("penempatan", penempatan)

		// Lanjutkan eksekusi jika token valid dan data berhasil diambil
		return c.Next()
	}
}

func CheckPermissions(allowedPosition []string, allowedDepartment []string, allowedPlacement []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Mengambil data pengguna dari lokal
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

func GenerateToken(user Domain.User) (string, error) {
	// Tentukan waktu kedaluwarsa token (contoh: 1 jam dari sekarang)
	expiredAt := time.Now().Add(time.Hour * 1).Unix()

	// Buat payload token dengan klaim waktu kedaluwarsa
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.UserID,
		"jabatan":    user.Jabatan,
		"departemen": user.Departemen,
		"penempatan": user.Penempatan,
		"exp":        expiredAt, // Klaim waktu kedaluwarsa
		// Anda dapat menambahkan informasi tambahan di sini
	})

	// Tandatangani token dengan kunci rahasia
	tokenString, err := claims.SignedString([]byte("secret")) // Ganti "secret" dengan kunci rahasia Anda yang lebih aman
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
