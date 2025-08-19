package password

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashAndCheckPassword(t *testing.T) {
	pass := "supersecret"

	hash, salt, err := HashPassword(pass, 16, bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(hash) == 0 {
		t.Fatalf("expected non-empty hash")
	}
	if len(salt) != 16 {
		t.Fatalf("expected salt length 16, got %d", len(salt))
	}

	// Check correct password
	if !CheckPasswordHash(pass, salt, hash) {
		t.Error("expected password to match")
	}

	// Check incorrect password
	if CheckPasswordHash("wrongpass", salt, hash) {
		t.Error("expected password NOT to match")
	}
}

func TestHashPassword_UsesDefaultCostWhenZero(t *testing.T) {
	pass := "secret"
	hash, salt, err := HashPassword(pass, 8, 0) // cost=0 -> should fallback to DefaultCost
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Ensure hash works
	if !CheckPasswordHash(pass, salt, hash) {
		t.Error("expected password to match with default cost")
	}
}

func TestAppendSaltToPass(t *testing.T) {
	pass := "mypassword"
	salt := []byte("abc")
	result := appendSaltToPass(pass, salt)

	expected := "mypassword:abc"
	if string(result) != expected {
		t.Errorf("expected %q, got %q", expected, string(result))
	}
}
