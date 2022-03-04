package helper

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

//generate hash from password
func GenerateHash(password string) string {
	hashpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error in generating password")
	}
	hashstring := fmt.Sprintf("%s", hashpass)
	return hashstring
}

//compare password and hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//get environment variables
func GetEnvVar(name string) string {
	err := godotenv.Load("D:/GolangPj/github.com/maoaeri/openapi/.env")
	if err != nil {
		fmt.Print(err)
		log.Fatalf("Error loading .env file")
	}
	value := os.Getenv(name)
	return value
}
