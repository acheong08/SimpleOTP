// Entries are the main data structure of the database. They are encrypted and stored
package database

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"errors"
	"io"
	"os"
)

type Entry struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Secret      string `json:"secret"`
}

// Marshals the JSON, encrypts it, and returns a base64 encoded string
func (e *Entry) Encrypt() string {
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
func Decrypt(entry string) (*Entry, error) {
	encryptedData, err := base64.StdEncoding.DecodeString(entry)
	if err != nil {
		return nil, errors.New("failed to decode base64")
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
		return nil, errors.New("invalid ciphertext")
	}

	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.New("failed to decrypt")
	}

	var entryData Entry
	err = json.Unmarshal(plaintext, &entryData)
	if err != nil {
		panic(err)
	}

	return &entryData, nil
}

type Entries struct {
	Entries map[string]string `json:"entries"` // Key is the hashed name, value is the base64 encoded encrypted entry
}

func (e *Entries) Search(name string) (Entry, error) {
	// Hash the name
	hashedName := sha512.Sum512([]byte(name))
	// Search the map for the hashed name
	// If it doesn't exist, return an error

	if entry, ok := e.Entries[string(hashedName[:])]; ok {
		decryptedEntry, err := Decrypt(entry)
		if err != nil {
			return Entry{}, err
		}
		return *decryptedEntry, nil
	} else {
		return Entry{}, errors.New("entry not found")
	}
}

func (e *Entries) Add(entry Entry) {
	// Hash the name
	hashedName := sha512.Sum512([]byte(entry.Name))
	// Add the entry to the map
	e.Entries[string(hashedName[:])] = entry.Encrypt()
}

func (e *Entries) Remove(name string) {
	// Hash the name
	hashedName := sha512.Sum512([]byte(name))
	// Remove the entry from the map
	delete(e.Entries, string(hashedName[:]))
}

func (e *Entries) List() ([]string, error) {
	names := make([]string, len(e.Entries))
	i := 0
	for entry := range e.Entries {
		decryptedEntry, err := Decrypt(entry)
		if err != nil {
			return nil, err
		}
		names[i] = decryptedEntry.Name
		i++
	}
	return names, nil

}

func (e *Entries) Save() error {
	// Gob encode the entries
	file, err := os.OpenFile("entries.gob", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(e)
	if err != nil {
		return err
	}
	return nil
}
