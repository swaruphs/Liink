package services

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"api.link/models"
	jwt "github.com/dgrijalva/jwt-go"
)

type keys struct {
	privateKey []byte
	publicKey  []byte
}

var keyInstance *keys = nil

func GenerateToken(userId int) (string, error) {

	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims["exp"] = time.Now().Add(time.Hour * time.Duration(72)).Unix()
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["userid"] = userId
	tokenString, err := token.SignedString(getPrivateKey())

	if err != nil {
		fmt.Println("something wrong with generating token string")
		return tokenString, err
	}

	return tokenString, nil
}

func RegisterUser(username string, password string) (bool, error) {
	rows, err := models.DB.Query("select username from users where username = $1", username)
	if err != nil {
		return false, err
	}

	if rows.Next() {
		return false, errors.New("username already exist")
	}

	_, ierr := models.DB.Exec("insert into users (username, password) values($1, $2)", username, password)
	if ierr != nil {
		return false, ierr
	}
	return true, nil
}

func Login(username string, password string) (map[string]string, error) {

	var token string
	result := map[string]string{}
	rows, err := models.DB.Query("select id from users where username = $1 and password = $2",
		username, password)

	if err != nil {
		return result, err
	}

	if rows.Next() == false {
		return result, errors.New("username and password not found")
	}

	var userID int
	rows.Scan(&userID)
	token, err = GenerateToken(userID)
	if err != nil {
		return result, err
	}
	result["token"] = token
	return result, err

}

func RefreshToken(userID int) (map[string]string, error) {

	result := map[string]string{}
	token, err := GenerateToken(userID)
	if err != nil {
		return result, err
	}

	result["token"] = token
	return result, err
}

//private helpers
func getPrivateKey() []byte {

	if keyInstance == nil {
		instantiateKeys()
	}

	return keyInstance.privateKey
}

func GetPublicKey() []byte {

	if keyInstance == nil {
		instantiateKeys()
	}
	return keyInstance.publicKey
}

func instantiateKeys() {
	privateKey, _ := ioutil.ReadFile("./settings/todolist.rsa")
	publicKey, _ := ioutil.ReadFile("./settings/todolist.rsa.pub")
	keyInstance = new(keys)
	keyInstance.privateKey = privateKey
	keyInstance.publicKey = publicKey
}
