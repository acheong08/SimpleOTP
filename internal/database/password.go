// Authentication
package database

import "crypto/sha256"

type Password struct {
	Key []byte `json:"key"` // Must be 32 bytes
}

// Stores the sha256 hash of the password to ensure 32 bytes
func SetPassword(pwd string) {
	hash := sha256.Sum256([]byte(pwd))
	// Check if the hash is 32 bytes
	if len(hash) != 32 {
		panic("Password hash is not 32 bytes")
	}
	key = Password{
		Key: hash[:],
	}
}

var key Password
