package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"reflect"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Function to generate a random salt
func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// Function to hash a password with Argon2id
func hashPassword(password string, salt []byte) string {
	// Define the parameters for Argon2id
	time := uint32(1)           // number of iterations
	memory := uint32(64 * 1024) // memory in KiB
	threads := uint8(4)         // number of threads
	keyLen := uint32(32)        // length of the generated key

	// Generate the hash
	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	// Encode the salt and hash to base64 for storage
	saltBase64 := base64.RawStdEncoding.EncodeToString(salt)
	hashBase64 := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf("%s$%s", saltBase64, hashBase64)
}

// Function to verify a password against a stored hash
func verifyPassword(password, storedHash string) bool {
	// Split the stored hash into its components (salt and hash)
	parts := strings.Split(storedHash, "$")
	if len(parts) != 2 {
		return false
	}
	saltBase64 := parts[0]
	hashBase64 := parts[1]

	// Decode the salt and hash from base64
	salt, err := base64.RawStdEncoding.DecodeString(saltBase64)
	if err != nil {
		return false
	}
	hash, err := base64.RawStdEncoding.DecodeString(hashBase64)
	if err != nil {
		return false
	}

	// Hash the provided password with the decoded salt
	computedHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	// Compare the computed hash with the stored hash
	return compareHashes(hash, computedHash)
}

// Function to compare two hashes for equality
func compareHashes(hash1, hash2 []byte) bool {
	if len(hash1) != len(hash2) {
		return false
	}
	for i := 0; i < len(hash1); i++ {
		if hash1[i] != hash2[i] {
			return false
		}
	}
	return true
}

func main() {
	password := "mysecurepassword"

	// Generate a random salt
	salt, err := generateSalt(16)
	if err != nil {
		log.Fatalf("Error generating salt: %v", err)
	}
	fmt.Println(reflect.TypeOf(salt))
	fmt.Println(salt)

	// Hash the password with the generated salt
	hashedPassword := hashPassword(password, salt)

	fmt.Printf("Hashed password: %s\n", hashedPassword)

	// Now let's verify the password
	inputPassword := "mysecurepassword"
	fmt.Println("---------------")
	if verifyPassword(inputPassword, hashedPassword) {
		fmt.Println("Password verified successfully!")
	} else {
		fmt.Println("Incorrect password!")
	}
}
