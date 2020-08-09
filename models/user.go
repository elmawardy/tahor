package models

import (
	"errors"
	"fmt"
	"strconv"

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

type UserBasicInfo struct {
	Name  string
	Email string
}

// func UserIDByJwt(jwt string) (u User, err error) {

// 	err = global.DB.First(&u).Where("email = ?", email).Error
// 	if err != nil {
// 		return u, err
// 	}
// 	return u, err
// }

func (u *User) InjectIdByJwt(jwt string) error {
	realJWT, err := global.Cachestore.HGet("user:"+u.Email, "jwt").Result()
	if err == redis.Nil {
		return errors.New("jwt not found")
	} else if err != nil {
		return err
	}

	if realJWT == jwt {
		id, err := global.Cachestore.HGet("user:"+u.Email, "id").Result()
		intID, _ := strconv.Atoi(id)
		u.ID = uint(intID)

		return err
	}

	return errors.New("jwt not the same")

}

func UserGetBasicInfo(jwt string) (info UserBasicInfo, err error) {
	u := User{}
	claims, err := ParseToken(jwt)
	if claims == nil {
		// log.Println(err)
		fmt.Println(`{"Status":"Error","Msg":"Unauthorized"}`)
		return
	}
	email := fmt.Sprintf("%v", claims["email"])
	global.DB.Where("email = ?", email).Find(&u)
	info.Email = u.Email
	info.Name = u.Name

	return
}

func UserLogout(tokenString string) error {
	claims, err := ParseToken(tokenString)
	if claims == nil {
		// log.Println(err)
		fmt.Println(`{"Status":"Error","Msg":"Unauthorized"}`)
		return errors.New("No Claims found in the token")
	}
	email := fmt.Sprintf("%v", claims["email"])
	err = global.Cachestore.Del("user:" + email).Err()

	return err
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

			userData := map[string]interface{}{
				"jwt": tokenString,
				"id":  u.ID,
			}
			err = global.Cachestore.HMSet("user:"+u.Email, userData).Err()
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

func ParseToken(userToken string) (jwt.MapClaims, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.

	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("my-token-key"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return claims, err
	}
}
