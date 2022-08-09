package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type AuthTokenClaims struct {
	jwt.StandardClaims        // Default Token Claims
	TokenUUID          string `json:"uuid"` // Token's UUID
}

const _tokenTTL = 3600 * time.Second // 토큰의 TTL, 1시간

var _secret = []byte("SECRET_FOR_YOUR_TOKEN") // SignedToken을 만들기 위한 secret
var _method = jwt.SigningMethodHS256          // HS256 알고리즘을 사용

// GenerateToken 생성한 UUID를 Claim으로 포함한 JWT 토큰을 생성하는 함수
func GenerateToken() (string, error) {
	now := time.Now()
	newUUID := uuid.NewString()

	claim := AuthTokenClaims{}
	claim.IssuedAt = now.Unix()
	claim.ExpiresAt = now.Add(_tokenTTL).Unix()
	claim.TokenUUID = newUUID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(_secret)
}

// ValidateToken 입력받은 토큰을 확인하는 함수
func ValidateToken(token string) (*AuthTokenClaims, error) {
	t, err := jwt.ParseWithClaims(token, &AuthTokenClaims{}, func(tk *jwt.Token) (interface{}, error) {
		if tk.Method.Alg() != _method.Alg() {
			err := errors.New("signing method mismatch")
			return nil, err
		}
		return _secret, nil
	})

	if err != nil {
		err = errors.New(err.Error())
		return nil, err
	}

	if claims, ok := t.Claims.(*AuthTokenClaims); ok && t.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
