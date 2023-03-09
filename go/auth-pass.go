package main

import (
	"fmt"

	"github.com/ermites-io/passwd"
)

func main7() {
	password := "my1337p4ssw0rd!"
	p, err := passwd.New(passwd.Argon2idDefault)
	if err != nil {
		// handle error
	}
	hashed, err := p.Hash([]byte(password))
	if err != nil {
		// handle error
	}
	fmt.Println(hashed)
	fmt.Println(string(hashed))

	err2 := passwd.Compare(hashed, []byte(password))
	if err2 != nil {
		// handle error
		fmt.Println("No match passwords")
	}
	//fmt.Printf("Match %s and %s\n", hashed, password)

	p2, err3 := passwd.NewMasked(passwd.Argon2idDefault)
	if err3 != nil {
		// handle error
	}

	// set the hashing key.
	if p2.SetKey([]byte("myhashingsecret")) != nil {
		// handle error
	}

	hashed2, err4 := p2.Hash([]byte("my1337p4ssw0rd!"))
	if err4 != nil {
		// handle error
	}

	fmt.Println(hashed2)
	fmt.Println(string(hashed2))

}
