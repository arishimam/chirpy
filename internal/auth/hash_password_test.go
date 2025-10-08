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
	password := "hello"

	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("Error calling HashPassword: %v", err)
	}

	match, err := CheckPasswordHash(password, hash)
	if err != nil {
		t.Errorf("Error calling CheckPasswordHash: %v", err)
	}

	if !match {
		t.Errorf("password and hashed password do not match!")
	}

	// negative case
	wrongPassword := "helloooo"

	match, _ = CheckPasswordHash(wrongPassword, hash)

	if match {
		t.Errorf("wrong password and hashed password should not match!")

	}

}
