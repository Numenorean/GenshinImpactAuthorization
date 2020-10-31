package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"log"
	"math/big"
)

var publicKey rsa.PublicKey

// Default encryption key in hex. Was extracted from the web authorization
const keyInHex string = "c3bde91d3cc1cddc06219bfbe4b494fe609afb708e4372c34aa9db31e43657d200ee585b888f377006eb6b2183cd9912751bcc9b0c817ba035b6784a66e6c31b2fdcecf44c5709dbeaae7e75a842dbaa3d17c6d3132296821c5488e743df3e94c557d5edfe19b2570a24a0e5c59401200a7f900a01ace766c5a1832dca2fb111"

func init() {
	publicKey.N, _ = new(big.Int).SetString(keyInHex, 16)
	publicKey.E = 65537
}

func encryptPassword(password []byte) string {
	// Password encryption with RSA 1024 Bit key and PKCS1v15 padding
	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, &publicKey, password)
	if err != nil {
		log.Panic(err)
	}
	return base64.StdEncoding.EncodeToString(encrypted)
}
