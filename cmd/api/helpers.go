package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
	"strings"
)

func (app *Config) sendEmail(lacubaMsg LacubaMessage) {
	app.Wait.Add(1)
	app.Mailer.MailerChan <- lacubaMsg
}

func deriveKey(passphrase string, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		// http://www.ietf.org/rfc/rfc2898.txt
		// Salt.
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(passphrase), salt, 1000, 32, sha256.New), salt
}

func encryptToken(passphrase, plaintext string) (string, error) {
	key, salt := deriveKey(passphrase, nil)
	iv := make([]byte, 12)
	// http://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38d.pdf
	// Section 8.2
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	data := aesgcm.Seal(nil, iv, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(salt) + "-" + hex.EncodeToString(iv) + "-" + hex.EncodeToString(data))), nil

}

func decryptToken(passphrase, ciphertext string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	arr := strings.Split(string(decoded), "-")
	salt, err := hex.DecodeString(arr[0])
	if err != nil {
		return "", err
	}

	iv, err := hex.DecodeString(arr[1])
	if err != nil {
		return "", err
	}

	data, err := hex.DecodeString(arr[2])
	if err != nil {
		return "", err
	}

	key, _ := deriveKey(passphrase, salt)

	b, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(b)
	if err != nil {
		return "", err
	}

	data, err = aesgcm.Open(nil, iv, data, nil)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
