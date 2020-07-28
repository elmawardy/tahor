package models

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/elmawardy/tahor/global"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique;not null"`
	Password string
}

func (u *User) Login() (tokenString string, err error) {

	plainTextPasswd := u.Password
	global.DB.Where("email = ?", u.Email).Find(&u)

	if CheckPasswordHash(plainTextPasswd, u.Password) {
		tokenString, err = global.Cachestore.HGet("user:"+u.Email, "jwt").Result()
		if err == redis.Nil {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"email": u.Email,
			})
			tokenString, err = token.SignedString([]byte("my-token-key"))
			if err != nil {
				return "", errors.New("Error generating jwt")
			}

			err = global.Cachestore.HSet("user:"+u.Email, "jwt", tokenString).Err()
			if err != nil {
				fmt.Println(err)
				return "", err
			}
		} else if err != nil {
			fmt.Println(err)
			return "", err
		}

		return tokenString, nil
	}

	return "", errors.New("Wrong password")

}

func (u *User) Register() error {
	u.Password, _ = HashPassword(u.Password)
	errs := global.DB.Create(u).GetErrors()

	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		return errors.New("Error creating user")
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
