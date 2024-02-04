package main

func main() {
	// This should come from outside
	key := []byte("your-32-byte-aes") // Replace with your key

	// Encrypt a file
	// These should be arguments
	if err := encrypt("./plaintext.txt", "./encrypted.bin", key); err != nil {
		panic(err)
	}
	println("Encryption successful")

	// Decrypt the file
	if err := decrypt("./encrypted.bin", "./decrypted.txt", key); err != nil {
		panic(err)
	}
	println("Decryption successful")
}
