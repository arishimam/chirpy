package auth

import (
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

}
