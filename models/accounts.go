package models

import (
	u "api-go-deploy-heroku/utils"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//JWT claims struct

type Token struct {
	UserId uint
	jwt.StandardClaims
}

// a struct to rep user account
type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

//Validate incoming user details...
func (account *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(account.Email, "@") {
		return u.Message("01", false, "Email address is required"), false
	}
	if len(account.Password) < 6 {
		return u.Message("01", false, "minimum lenght of password is 6"), false
	}

	//Email must be unique
	temp := &Account{}

	//check for errors and duplicate emails
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message("01", false, "Connection Error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message("01", false, "Email address already in use by another user."), false
	}
	return u.Message("01", false, "Requirement passed"), true
}

func (account *Account) Create() map[string]interface{} {
	if resp, ok := account.Validate(); !ok {
		return resp
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID < 1 {
		return u.Message("01", false, "Failed to create account, connection error.")
	}

	//Create new JWT token for the newly registered account
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //delete password
	response := u.Message("00", true, "Account has been created")
	response["account"] = account
	return response
}

func Login(email, password string) map[string]interface{} {
	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message("01", false, "Email address not found")
		}
		return u.Message("01", false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //invalid password
		return u.Message("01", false, "Invalid login credentials. Please try again")
	} //working! Logged In
	account.Password = ""

	//Create JWT Token
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //store the token in the response

	resp := u.Message("00", true, "Logged in")
	resp["account"] = account
	return resp
}

func GetUser(u uint) *Account {
	acc := &Account{}
	GetDB().Table("account").Where("id = ?", u).First(acc)
	if acc.Email == "" {
		return nil
	}
	acc.Password = ""
	return acc
}
