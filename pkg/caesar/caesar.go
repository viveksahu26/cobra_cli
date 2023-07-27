package caesar

import (
	"strings"
	"unicode"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

// encrypt gets the cipher character present at new index.
// The new index is obtained by rotating the index of plaintext character by rotation value.
func Encrypt(plaintext string, rotations int) string {
	var ciphertext string

	for _, character := range plaintext {
		index := getCharacterIndex(strings.ToLower(string(character)))
		if index == -1 {
			ciphertext += string(character)
		} else {
			cipher := alphabets[getCipherIndex(index, rotations)]
			if unicode.IsUpper(character) {
				ciphertext += strings.ToUpper(string(cipher))
			} else {
				ciphertext += string(cipher)
			}
		}
	}
	return ciphertext
}

func Decrypt(ciphertext string, rotations int) string {
	var plaintext string
	for _, character := range ciphertext {
		index := getCharacterIndex(strings.ToLower(string(character)))
		if index == -1 {
			plaintext += string(character)
		} else {
			plainCharacter := alphabets[getPlaintextIndex(index, int(rotations))]
			if unicode.IsUpper(character) {
				plaintext += strings.ToUpper(string(plainCharacter))
			} else {
				plaintext += string(plainCharacter)
			}
		}
	}
	return plaintext
}

// get plaintext character or ciphertext character index from alphabets list
func getCharacterIndex(element string) int {
	return strings.Index(alphabets, element)
}

// CipherIndex: get new index by rotating the plaintext
// character index by rotation value from alphabets list
func getCipherIndex(plainIndex int, key int) int {
	newIndex := plainIndex + key
	if newIndex > 25 {
		return newIndex % 26
	}
	return newIndex
}

// PlainText Index: get new index by reverse-rotating the cipher
// character index by rotation value from alphabets list
func getPlaintextIndex(cipherIndex int, key int) int {
	newIndex := cipherIndex - key
	if newIndex < 0 {
		newIndex = 26 + (newIndex % 26)
	}
	if newIndex > 25 {
		return newIndex % 26
	}
	return newIndex
}
