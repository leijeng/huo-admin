package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var (
	TokenInvalid = errors.New("Couldn't handle this token:")
)

type TypJWT struct {
	SigningKey []byte
}

type CustomTymonClaims struct {
	Subject  int64  `json:"sub"`
	Prv      string `json:"prv"`
	UserId   int    `json:"uid,omitempty"`
	RoleId   int    `json:"rid,omitempty"`
	Phone    string `json:"mob,omitempty"`
	Nickname string `json:"nick,omitempty"`
	Username string `json:"username,omitempty"`
	IsLogin  int    `json:"isLogin,omitempty"`
	jwt.StandardClaims
}

func NewJWT(secret string) *TypJWT {
	return &TypJWT{SigningKey: []byte(secret)}
}

// 创建token
func (j *TypJWT) CreateTymonToken(claims CustomTymonClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	res, err := token.SignedString(j.SigningKey)
	fmt.Println("err:", err)
	return res, err
}

// 解析token
func (j *TypJWT) ParseTymonToken(tokenString string) (*CustomTymonClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomTymonClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Panicln("unexpected signing method")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.SigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomTymonClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *TypJWT) RefreshTymonToken(tokenString string) (string, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomTymonClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	fmt.Println("RefreshTymonToken err  ======", err)

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*CustomTymonClaims)

	fmt.Println("ok ======", ok)
	fmt.Println("Valid ======", token.Valid)

	if ok && token.Valid {

		claims.StandardClaims.ExpiresAt = time.Now().Add(7 * 24 * time.Hour).Unix()
		return j.CreateTymonToken(*claims)
	}

	return "", TokenInvalid
}

func (j *TypJWT) GenerateTymonToken(claims CustomTymonClaims) (string, error) {

	token, err := j.CreateTymonToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}
