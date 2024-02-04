package main

import (
	"fmt"
	"os"
)

func main() {
	// This should come from outside
	env := os.Getenv("ENCRYPTER_KEY")

	if env == "" {
		fmt.Println("Environment variable ENCRYPTER_KEY is not set.")
		return
	}

	if len(os.Args) < 4 {
		fmt.Println("Usage: crypter <operation> <input_file> <output_file>")
		return
	}

	operation := os.Args[1]
	inputFile := os.Args[2]
	outputFile := os.Args[3]

	key := []byte(env)

	switch operation {
	case "encrypt":
		if err := encrypt(inputFile, outputFile, key); err != nil {
			panic(err)
		}
		println("Action successful")
	case "decrypt":
		if err := decrypt(inputFile, outputFile, key); err != nil {
			panic(err)
		}
		println("Action successful")
	}
}
