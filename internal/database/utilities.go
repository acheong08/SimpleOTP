package database

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"os"
)

// Marshals the JSON, encrypts it, and returns a base64 encoded string
func Encrypt(e any) string {
	plaintext, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(key.Key)
	if err != nil {
		panic(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)
	encryptedData := append(nonce, ciphertext...)
	return base64.StdEncoding.EncodeToString(encryptedData)
}

// Decrypts a base64 encoded and encrypted string, and unmarshals it
func Decrypt(data string, obj any) error {
	encryptedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return errors.New("failed to decode base64")
	}

	block, err := aes.NewCipher(key.Key)
	if err != nil {
		panic(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonceSize := aesGCM.NonceSize()
	if len(encryptedData) < nonceSize {
		return errors.New("invalid ciphertext")
	}

	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return errors.New("failed to decrypt")
	}

	err = json.Unmarshal(plaintext, obj)
	if err != nil {
		panic(err)
	}

	return nil
}

func hash(str string) string {
	hashedName := sha512.Sum512([]byte(str))
	// Hex
	return hex.EncodeToString(hashedName[:])
}

func SaveFile(obj any, filePath string) error {
	// Gob encode the entries
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(obj)
	if err != nil {
		return err
	}
	return nil
}
func LoadFile(obj any, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(obj)
	if err != nil {
		return err
	}
	return nil
}
