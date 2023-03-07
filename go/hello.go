package main

import (
	"fmt"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func main2() {
	fmt.Println("hello")
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v", s, s)
	}
	fmt.Printf("\n")
	b := "true"
	if s, err := strconv.ParseBool(b); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	f := "3.1415926535"
	if s, err := strconv.ParseFloat(f, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	fmt.Println(strings.ToLower("Gopher"))
	fmt.Println(strings.ToTitle("dünyanın ilk borsa yapısı Aizonai kabul edilir"))
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))

	var re = regexp.MustCompile(`(?m)^(?P<name>[a-zA-Z0-9.!#$%&'*+/=?^_ \x60{|}~-]+)@(?P<domain>[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*)$`)
	var str = `// https://html.spec.whatwg.org/#valid-e-mail-address
me@example.com
me@example.co.uk
me@example.new
bad-example`
	for i, match := range re.FindAllString(str, -1) {
		fmt.Println(match, "found at index", i)
	}

	for _, email := range []string{
		"good@exmaple.com",
		"bad-example",
	} {
		fmt.Printf("%18s valid: %t\n", email, validEmail(email))
	}

}
