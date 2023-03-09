package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const saltSize = 32

func main6() {
	email := []byte("john.doe@somedomain.com")
	password := []byte("47;u5:B(95m72;Xq")
	// create random word
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	// let's create SHA256(salt+password)
	hash := sha256.New()
	hash.Write(salt)
	hash.Write(password)
	h := hash.Sum(nil)
	fmt.Printf("email : %s\n", string(email))
	fmt.Printf("password: %s\n", string(password))
	fmt.Printf("salt : %x\n", salt)
	fmt.Printf("hash : %x\n", h)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	fmt.Printf("email : %s\n", string(email))
	fmt.Printf("password : %s\n", string(password))
	fmt.Printf("hashed password: %x\n", hashedPassword)

	// AUTENTICACION
	fmt.Printf("Comparing: %s and %s\n", string(hashedPassword), string(password))
	if bcrypt.CompareHashAndPassword(hashedPassword, password) != nil {
		// passwords do not match
		// passwords mismatch should be logged (see Error Handling and Logging) // error should be returned so that a GENERIC message "Sign-in attempt has // failed, please check your credentials" can be shown to the user.
		fmt.Println("Error with username or password")
	} else {
		fmt.Println("Password match")

	}

}
