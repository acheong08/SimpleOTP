// Authentication
package database

import (
	"crypto/sha256"

	"github.com/acheong08/SimpleOTP/internal/utilities"
)

type Password struct {
	Key [32]byte `json:"key"` // Must be 32 bytes
}

// Stores the sha256 hash of the password to ensure 32 bytes
func SetPassword(pwd string) string {
	hash := sha256.Sum256([]byte(pwd))
	// Check if the hash is 32 bytes
	if len(hash) != 32 {
		panic("Password hash is not 32 bytes")
	}
	key = Password{
		Key: hash,
	}
	if password_hash_loaded {
		// Compare the hashes
		if utilities.Hash(pwd) != FileStore.PasswordHash.Hash {
			return "failed"
		}
	} else {
		// Save the hash
		FileStore.PasswordHash.Hash = utilities.Hash(pwd)
		password_hash_loaded = true
	}
	return "success"
}

// Salt and hash are hex encoded
type PasswordHash struct {
	Hash string `json:"hash"`
}

var key Password
var password_hash_loaded bool = false

func init() {
	// Load the password hash
	err := FileStore.Load()
	password_hash_loaded = err == nil
}
