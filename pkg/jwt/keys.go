package jwt

import (
	"hinsun-backend/pkg/security"
	"os"
)

type KeyManager struct {
	algorithm  security.Algorithm
	accessKey  *security.KeyPair
	refreshKey *security.KeyPair
}

func NewKeyManager(alg string, keySize int) *KeyManager {
	cryptography := security.NewCryptography()
	rootPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	algorithm, err := security.AlgorithmFromString(alg)
	if err != nil {
		panic(err)
	}

	keyPath := rootPath + "/keys"
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		err := os.Mkdir(keyPath, 0700)
		if err != nil {
			panic(err)
		}
	}

	accessPath := keyPath + "/access"
	refreshPath := keyPath + "/refresh"

	accessKey, err := cryptography.InitializeKeyPair(accessPath, algorithm, keySize)
	if err != nil {
		panic(err)
	}

	refreshKey, err := cryptography.InitializeKeyPair(refreshPath, algorithm, keySize)
	if err != nil {
		panic(err)
	}

	return &KeyManager{
		algorithm:  algorithm,
		accessKey:  accessKey,
		refreshKey: refreshKey,
	}
}
