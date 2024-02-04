// encrypt_test.go
package main

import (
	"crypto/rand"
	"os"
	"testing"
)

// TestEncryptDecrypt is a test function for encrypt and decrypt functions
func TestEncryptDecrypt(t *testing.T) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		t.Fatal("Error generating random key:", err)
	}

	// Create a temporary test file
	testFile := "testfile.txt"
	if err := os.WriteFile(testFile, []byte("Hello, world!"), 0644); err != nil {
		t.Fatal("Error creating test file:", err)
	}
	defer os.Remove(testFile) // Defer the file removal

	// Encrypt and then decrypt the test file
	encryptedFile := "encrypted.bin"
	decryptedFile := "decrypted.txt"
	if err := encrypt(testFile, encryptedFile, key); err != nil {
		t.Fatal("Error encrypting file:", err)
	}
	defer os.Remove(encryptedFile) // Defer the file removal

	if err := decrypt(encryptedFile, decryptedFile, key); err != nil {
		t.Fatal("Error decrypting file:", err)
	}

	// Read the decrypted file and verify its content
	decryptedContent, err := os.ReadFile(decryptedFile)
	if err != nil {
		t.Fatal("Error reading decrypted file:", err)
	}

	expectedContent := "Hello, world!"
	if string(decryptedContent) != expectedContent {
		t.Fatalf("Expected decrypted content to be %s, got %s", expectedContent, decryptedContent)
	}
}
