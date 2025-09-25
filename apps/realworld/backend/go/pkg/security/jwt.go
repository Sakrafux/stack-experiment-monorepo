package security

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenData struct {
	Id   int64
	Role string
}

func CreateAccessToken(data *TokenData, accessSecret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": fmt.Sprintf("%d", data.Id),
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
		"iat":     time.Now().Unix(),
		"role":    data.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(accessSecret))
	return tokenString, err
}

func CreateRefreshToken(data *TokenData, refreshSecret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": fmt.Sprintf("%d", data.Id),
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
		"iat":     time.Now().Unix(),
		"role":    data.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(refreshSecret))
}

func ValidateAccessToken(tokenStr, accessSecret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(accessSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid access token")
	}

	return token, nil
}

func ValidateRefreshToken(tokenStr, refreshSecret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(refreshSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid refresh token")
	}

	return token, nil
}

func ExtractTokenData(token *jwt.Token) (*TokenData, error) {
	claims := token.Claims.(jwt.MapClaims)
	userId, err := strconv.ParseInt(claims["user_id"].(string), 10, 64)
	if err != nil {
		return nil, err
	}
	role, ok := claims["role"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid access token")
	}
	return &TokenData{Id: userId, Role: role}, nil
}
