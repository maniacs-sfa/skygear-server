package oddb

import (
	"bytes"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestNewUserInfo(t *testing.T) {
	info := NewUserInfo("userinfoid", "john.doe@example.com", "secret")

	if info.ID != "userinfoid" {
		t.Fatalf("got info.ID = %v, want userinfoid", info.ID)
	}

	if info.Email != "john.doe@example.com" {
		t.Fatalf("got info.Email = %v, want john.doe@example.com", info.Email)
	}

	if bytes.Equal(info.HashedPassword, nil) {
		t.Fatalf("got info.HashPassword = %v, want non-empty value", info.HashedPassword)
	}
}

func TestNewUserInfoWithEmptyID(t *testing.T) {
	info := NewUserInfo("", "jane.doe@example.com", "anothersecret")

	if info.ID == "" {
		t.Fatalf("got empty info.ID, want non-empty string")
	}

	if info.Email != "jane.doe@example.com" {
		t.Fatalf("got info.Email = %v, want jane.doe@example.com", info.Email)
	}

	if bytes.Equal(info.HashedPassword, nil) {
		t.Fatalf("got info.HashPassword = %v, want non-empty value", info.HashedPassword)
	}
}

func TestNewAnonymousUserInfo(t *testing.T) {
	info := NewAnonymousUserInfo()
	if info.ID == "" {
		t.Fatalf("got info.ID = %v, want \"\"", info.ID)
	}

	if info.Email != "" {
		t.Fatalf("got info.Email = %v, want empty string", info.Email)
	}

	if len(info.HashedPassword) != 0 {
		t.Fatalf("got info.HashPassword = %v, want zero-length bytes", info.HashedPassword)
	}
}

func TestSetPassword(t *testing.T) {
	info := UserInfo{}
	info.SetPassword("secret")
	err := bcrypt.CompareHashAndPassword(info.HashedPassword, []byte("secret"))
	if err != nil {
		t.Fatalf("got err = %v, want nil", err)
	}
}

func TestIsSamePassword(t *testing.T) {
	info := UserInfo{}
	info.SetPassword("secret")
	if !info.IsSamePassword("secret") {
		t.Fatalf("got UserInfo.HashedPassword = %v, want a hashed \"secret\"", info.HashedPassword)
	}
}
