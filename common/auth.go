package common

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
	Username string `json:"username"`
    Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(username, role string) (string, error) {
    // Xác định thời gian hết hạn của token (12 giờ kể từ thời điểm hiện tại)
    expirationTime := time.Now().Add(12 * time.Hour)
    
    // Tạo đối tượng Claims chứa thông tin người dùng và thời gian hết hạn của token
    claims := &Claims{
        Username: username,
        Role: role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime), // Gán thời gian hết hạn cho token
        },
    }
    
    // Tạo token mới với claims và phương thức ký HS256 (HMAC SHA-256)
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
    // Ký token bằng khóa bí mật (jwtKey)
    tokenString, err := token.SignedString(JwtKey)
    
    // Trả về token đã ký và lỗi (nếu có)
    return tokenString, err
}
