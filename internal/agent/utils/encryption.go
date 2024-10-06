package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"os"
)

func InitPublicKey(filePath string) (*rsa.PublicKey, error) {
	publicKeyPEM, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	publicKeyBlock, _ := pem.Decode(publicKeyPEM)
	key, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}
	var publicKey *rsa.PublicKey
	switch v := key.(type) {
	case *rsa.PublicKey:
		publicKey = v
	default:
		return nil, errors.New("unexpected key type")
	}
	return publicKey, nil
}

func EncryptString(publicKey *rsa.PublicKey, plaintext []byte) (string, error) {
	msgLen := len(plaintext)
	// Не понял пока как подобрать число чтобы не было ошибки crypto/rsa: message too long for RSA key size.
	step := publicKey.Size() - 15 //nolint:mnd // не понял пока как подобрать число
	var encryptedBytes []byte

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}
		encryptedBlockBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext[start:finish])
		if err != nil {
			return "", err
		}

		encryptedBytes = append(encryptedBytes, encryptedBlockBytes...)
	}
	return hex.EncodeToString(encryptedBytes), nil
}

func InitPrivateKey(filePath string) (*rsa.PrivateKey, error) {
	privateKeyPEM, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	privateKeyBlock, _ := pem.Decode(privateKeyPEM)
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func DecryptString(privateKey *rsa.PrivateKey, ciphertext string) (string, error) {
	ciphertextBytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	msgLen := len(ciphertextBytes)
	var decryptedBytes []byte
	step := privateKey.PublicKey.Size()

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}
		decryptedBlockBytes, errDecrypt := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertextBytes[start:finish])
		if errDecrypt != nil {
			return "", errDecrypt
		}
		decryptedBytes = append(decryptedBytes, decryptedBlockBytes...)
	}
	return string(decryptedBytes), nil
}
