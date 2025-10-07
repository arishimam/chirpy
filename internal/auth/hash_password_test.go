package auth

import (
	"fmt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "tester"
	hashed_password, err := HashPassword(password)
	if err != nil {
		t.Errorf("Error occured when called HashPassword function")
	}

	if hashed_password == "" || hashed_password == "tester" {
		t.Errorf("Password was not successfully hashed!\nHashed password = %v", hashed_password)
	}

}

func TestCheckPasswordHash(t *testing.T) {
	test := "hello"
	pass1, err := HashPassword(test)
	if err != nil {
		t.Errorf("Error occured when calling HashPassword function")
	}

	pass2, err := HashPassword(test)
	if err != nil {
		t.Errorf("Error occured when calling HashPassword function")
	}

	if pass1 == pass2 {
		t.Errorf("pass1 and pass2 should be hashed to different passwords")
	}
	//fmt.Println("pass1 ", pass1)
	//fmt.Println("pass2 ", pass2)

}
