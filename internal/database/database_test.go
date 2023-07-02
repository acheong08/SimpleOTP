package database_test

import (
	"crypto/sha256"
	"testing"

	"github.com/acheong08/SimpleOTP/internal/database"
	"github.com/acheong08/SimpleOTP/internal/utilities"
)

func TestEntryEncryptDecrypt(t *testing.T) {

	key := sha256.Sum256([]byte("somepassword"))

	// Create a test entry
	entry := &database.Entry{
		Name:        "Test Entry",
		Description: "This is a test entry",
		URL:         "https://example.com",
		Secret:      "testsecret",
	}

	// Encrypt the entry
	encrypted := utilities.Encrypt(entry, key[:])

	// Decrypt the encrypted entry
	var decrypted database.Entry = database.Entry{}
	err := utilities.Decrypt(encrypted, key[:], &decrypted)
	if err != nil {
		t.Errorf("Failed to decrypt entry: %v", err)
	}

	// Compare the decrypted entry with the original entry
	if decrypted.Name != entry.Name {
		t.Errorf("Name mismatch: expected %s, got %s", entry.Name, decrypted.Name)
	}
	if decrypted.Description != entry.Description {
		t.Errorf("Description mismatch: expected %s, got %s", entry.Description, decrypted.Description)
	}
	if decrypted.URL != entry.URL {
		t.Errorf("URL mismatch: expected %s, got %s", entry.URL, decrypted.URL)
	}
	if decrypted.Secret != entry.Secret {
		t.Errorf("Secret mismatch: expected %s, got %s", entry.Secret, decrypted.Secret)
	}
}

func TestEntriesAddGetRemove(t *testing.T) {
	database.SetPassword("01234567890123456789012345678901")

	// Create a test entry
	entry := database.Entry{
		Name:        "Test Entry",
		Description: "This is a test entry",
		URL:         "https://example.com",
		Secret:      "testsecret",
	}

	// Initialize the entries
	entries := database.Entries{
		Entries: make(map[string]string),
	}

	// Add the entry
	entries.Add(entry)

	// Get the entry
	retrievedEntry, err := entries.Get(entry.Name)
	if err != nil {
		t.Errorf("Failed to retrieve entry: %v", err)
	}

	// Compare the retrieved entry with the original entry
	if retrievedEntry.Name != entry.Name {
		t.Errorf("Name mismatch: expected %s, got %s", entry.Name, retrievedEntry.Name)
	}
	if retrievedEntry.Description != entry.Description {
		t.Errorf("Description mismatch: expected %s, got %s", entry.Description, retrievedEntry.Description)
	}
	if retrievedEntry.URL != entry.URL {
		t.Errorf("URL mismatch: expected %s, got %s", entry.URL, retrievedEntry.URL)
	}
	if retrievedEntry.Secret != entry.Secret {
		t.Errorf("Secret mismatch: expected %s, got %s", entry.Secret, retrievedEntry.Secret)
	}

	// Remove the entry
	entries.Remove(entry.Name)

	// Verify that the entry was removed
	_, err = entries.Get(entry.Name)
	if err == nil {
		t.Errorf("Entry was not removed")
	}
}

func TestEntriesSearch(t *testing.T) {
	database.SetPassword("01234567890123456789012345678901")

	// Create some test entries
	entries := database.Entries{
		Entries: make(map[string]string),
	}

	entry1 := database.Entry{
		Name:        "Test Entry 1",
		Description: "This is test entry number 1",
		URL:         "https://example1.com",
		Secret:      "testsecret1",
	}
	entry2 := database.Entry{
		Name:        "Test Entry 2",
		Description: "This is test entry number 2",
		URL:         "https://example2.com",
		Secret:      "testsecret2",
	}
	entry3 := database.Entry{
		Name:        "Another Entry",
		Description: "This is another test entry",
		URL:         "https://anotherexample.com",
		Secret:      "anothersecret",
	}

	// Add the entries
	entries.Add(entry1)
	entries.Add(entry2)
	entries.Add(entry3)

	// Search for entries with "test" in the name
	results, err := entries.Search("test")
	if err != nil {
		t.Errorf("Search error: %v", err)
	}

	// Verify that the correct entries were found
	if len(results) != 2 {
		t.Errorf("Unexpected number of search results. Expected 2, got %d", len(results))
	}
	if !contains(results, entry1.Name) {
		t.Errorf("Search result does not contain expected entry: %s", entry1.Name)
	}
	if !contains(results, entry2.Name) {
		t.Errorf("Search result does not contain expected entry: %s", entry2.Name)
	}
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func TestEntriesList(t *testing.T) {
	database.SetPassword("01234567890123456789012345678901")

	// Create some test entries
	entries := database.Entries{
		Entries: make(map[string]string),
	}

	entry1 := database.Entry{
		Name:        "Test Entry 1",
		Description: "This is test entry number 1",
		URL:         "https://example1.com",
		Secret:      "testsecret1",
	}
	entry2 := database.Entry{
		Name:        "Test Entry 2",
		Description: "This is test entry number 2",
		URL:         "https://example2.com",
		Secret:      "testsecret2",
	}

	// Add the entries
	entries.Add(entry1)
	entries.Add(entry2)

	// Get the list of entry names
	names, err := entries.List()
	if err != nil {
		t.Errorf("Failed to get entry list: %v", err)
	}

	// Verify that the correct entry names are present
	expectedNames := []string{entry1.Name, entry2.Name}
	if len(names) != len(expectedNames) {
		t.Errorf("Unexpected number of entries. Expected %d, got %d", len(expectedNames), len(names))
	}
	for _, name := range expectedNames {
		if !contains(names, name) {
			t.Errorf("Entry list does not contain expected entry: %s", name)
		}
	}
}
