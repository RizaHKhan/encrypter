package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

func encrypt(fileIn, fileOut string, key []byte) error {
	// Read the plaintext file
	plaintext, err := os.ReadFile(fileIn)
	if err != nil {
		return err
	}

	// Generate a random nonce
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	// Create a new AES cipher block using the provided key
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Create a new GCM cipher
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Encrypt the plaintext using AES-GCM
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	// Write the nonce and ciphertext to the output file
	ciphertext = append(nonce, ciphertext...)
	if err := os.WriteFile(fileOut, ciphertext, 0644); err != nil {
		return err
	}

	return nil
}

func decrypt(fileIn, fileOut string, key []byte) error {
	// Read the encrypted file
	ciphertext, err := os.ReadFile(fileIn)
	if err != nil {
		return err
	}

	// Extract the nonce from the ciphertext
	nonce := ciphertext[:12]
	ciphertext = ciphertext[12:]

	// Create a new AES cipher block using the provided key
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Create a new GCM cipher
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Decrypt the ciphertext using AES-GCM
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	// Write the decrypted plaintext to the output file
	if err := os.WriteFile(fileOut, plaintext, 0644); err != nil {
		return err
	}

	return nil
}
