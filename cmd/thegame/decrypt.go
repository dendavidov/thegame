package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
)

// Author Nic Raboy
// https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))

	return hex.EncodeToString(hasher.Sum(nil))
}

func Decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err.Error())
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)

	if err != nil {
		panic(err.Error())
	}

	return plaintext
}
