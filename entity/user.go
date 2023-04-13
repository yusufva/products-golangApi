package entity

import (
	"strings"
	"time"
	"tugas-sesi12/pkg/errrs"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userLevel string

const (
	Admin    userLevel = "admin"
	Customer userLevel = "customer"
)

var secret_key = "ApA HaYO"

var invalidTokenErr = errrs.NewUnauthenticatedError("invalid token")

type User struct {
	Id         int       `gorm:"primaryKey;not null" json:"id"`
	Email      string    `gorm:"unique;not null;type:varchar(255)" json:"email"`
	Password   string    `gorm:"type:text[];not null" json:"password"`
	Level      userLevel `gorm:"not null;default:customer" json:"level"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

func (u *User) parseToken(tokenString string) (*jwt.Token, errrs.MessageErr) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidTokenErr
		}

		return []byte(secret_key), nil
	})

	if err != nil {
		return nil, invalidTokenErr
	}

	return token, nil
}

func (u *User) bindTokenToUserEntity(claim jwt.MapClaims) errrs.MessageErr {
	if id, ok := claim["id"].(float64); !ok {
		return invalidTokenErr
	} else {
		u.Id = int(id)
	}

	if email, ok := claim["email"].(string); !ok {
		return invalidTokenErr
	} else {
		u.Email = email
	}

	if level, ok := claim["level"].(string); !ok {
		return invalidTokenErr
	} else {
		u.Level = userLevel(level)
	}

	return nil
}

func (u *User) ValidateToken(bearerToken string) errrs.MessageErr {
	isBearer := strings.HasPrefix(bearerToken, "Bearer")

	if !isBearer {
		return invalidTokenErr
	}

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		return invalidTokenErr
	}

	tokenString := splitToken[2]

	token, err := u.parseToken(tokenString)

	if err != nil {
		return err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return invalidTokenErr
	} else {
		mapClaims = claims
	}

	err = u.bindTokenToUserEntity(mapClaims)
	return err
}

func (u *User) tokenClaim() jwt.MapClaims {
	return jwt.MapClaims{
		"id":    u.Id,
		"email": u.Email,
		"level": u.Level,
		"exp":   time.Now().Add(time.Minute * 120).Unix(),
	}
}

func (u *User) signToken(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString([]byte(secret_key))

	return tokenString
}

func (u *User) GenerateToken() string {
	claims := u.tokenClaim()

	return u.signToken(claims)
}

func (u *User) HashPassword() errrs.MessageErr {
	salt := 8

	userPassword := []byte(u.Password)

	bs, err := bcrypt.GenerateFromPassword(userPassword, salt)

	if err != nil {
		return errrs.NewInternalServerError("something went wrong")
	}

	u.Password = string(bs)

	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
