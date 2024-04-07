package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
	"os"
	"path/filepath"
	"time"
)

var JWTsecret = []byte("ABAB")

type Claims struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
	jwt.StandardClaims
}

func DigestPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func CheckPassword(digestPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(digestPassword), []byte(password))
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GenerateAccessToken(id uint, username string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(1 * time.Hour)
	claims := Claims{
		ID:       id,
		UserName: username,
		PassWord: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "videoweb",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accassToken, err := tokenClaims.SignedString(JWTsecret)
	return accassToken, err
}

func GenerateRefreshToken(id uint, username string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(48 * time.Hour)
	claims := Claims{
		ID:       id,
		UserName: username,
		PassWord: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "videoweb",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accassToken, err := tokenClaims.SignedString(JWTsecret)
	return accassToken, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return JWTsecret, nil
		})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func CreateFolder(username string) error {
	avatarfilepath := "./upload/avatar/" + username
	err1 := os.Mkdir(avatarfilepath, 0755)
	if err1 != nil {
		fmt.Println(err1)
		return err1
	}
	videofilepath := "./upload/video/" + username
	err2 := os.Mkdir(videofilepath, 0755)
	if err2 != nil {
		fmt.Println(err2)
		return err2
	}
	return nil
}

func ParseAvatarExt(filename string) error {
	fileExt := filepath.Ext(filename)
	if fileExt != ".jpeg" && fileExt != ".jpg" && fileExt != ".png" {
		return errors.New("头像文件格式不正确")
	}
	return nil
}

func ParseVideoExt(filename string) bool {
	fileExt := filepath.Ext(filename)
	if fileExt != ".mp4" {
		return true
	}
	return false
}

func CheckCode(secret, code string) error {
	if totp.Validate(code, secret) {
		return nil
	} else {
		return errors.New("MFA验证失败")
	}
}
