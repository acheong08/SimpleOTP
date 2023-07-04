// Entries are the main data structure of the database. They are encrypted and stored
package database

import (
	"errors"
	"strings"

	"github.com/acheong08/SimpleOTP/internal/utilities"
)

type Entry struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Secret      string `json:"secret"`
}

type Entries struct {
	Entries map[string]string `json:"entries"` // Key is the hashed name, value is the base64 encoded encrypted entry
}

func (e *Entries) Get(name string) (*Entry, error) {
	// Hash the name
	hashedName := utilities.Hash(name)
	// Get the entry from the map
	encryptedEntry, ok := e.Entries[string(hashedName[:])]
	if !ok {
		return nil, errors.New("entry not found")
	}
	// Decrypt the entry
	var decryptedEntry Entry = Entry{}
	err := utilities.Decrypt(encryptedEntry, key.Key[:], &decryptedEntry)
	if err != nil {
		return nil, err
	}
	return &decryptedEntry, nil
}

func (e *Entries) Search(name string) ([]Entry, error) {
	names, err := e.List()
	if err != nil {
		return nil, err
	}
	var matches []Entry
	for _, n := range names {
		// Check if name is a substring of n
		if strings.Contains(strings.ToLower(n.Name), strings.ToLower(name)) ||
			strings.Contains(strings.ToLower(n.Description), strings.ToLower(name)) {
			matches = append(matches, n)
		}
	}
	return matches, nil
}

func (e *Entries) Add(entry Entry) {
	// Hash the name
	hashedName := utilities.Hash(entry.Name)
	// Add the entry to the map
	e.Entries[string(hashedName[:])] = utilities.Encrypt(&entry, key.Key[:])

}

func (e *Entries) Remove(name string) {
	// Hash the name
	hashedName := utilities.Hash(name)
	// Remove the entry from the map
	delete(e.Entries, string(hashedName[:]))
}

func (e *Entries) List() ([]Entry, error) {
	entries := make([]Entry, len(e.Entries))
	i := 0
	for _, entry := range e.Entries {
		var decryptedEntry Entry = Entry{}
		err := utilities.Decrypt(entry, key.Key[:], &decryptedEntry)
		if err != nil {
			return nil, err
		}
		entries[i] = decryptedEntry
		i++
	}
	return entries, nil

}
