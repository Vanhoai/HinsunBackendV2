package security

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"hinsun-backend/internal/core/log"
	"os"
	"path/filepath"
)

var (
	ErrGenerateKeyPair  = errors.New("failed to generate key pair")
	ErrInvalidKeyFormat = errors.New("invalid key format")
	ErrUnsupportedKey   = errors.New("unsupported key type")
)

type Cryptography interface {
	InitializeKeyPair(
		rootPath string,
		algorithm Algorithm,
		keySize int,
	) (*KeyPair, error)
	LoadPrivateKey(path string) (any, error)
	LoadPublicKey(path string) (any, error)
}

type cryptography struct{}

func NewCryptography() Cryptography {
	return &cryptography{}
}

func (c *cryptography) InitializeKeyPair(
	rootPath string,
	algorithm Algorithm,
	keySize int,
) (*KeyPair, error) {
	err := c.generateKeyPair(rootPath, algorithm, keySize)
	if err != nil {
		return nil, err
	}

	privateKey, err := c.LoadPrivateKey(filepath.Join(rootPath, "private_key.pem"))
	if err != nil {
		return nil, fmt.Errorf("failed to load private key: %w", err)
	}

	publicKey, err := c.LoadPublicKey(filepath.Join(rootPath, "public_key.pem"))
	if err != nil {
		return nil, fmt.Errorf("failed to load public key: %w", err)
	}

	return &KeyPair{
		SigningKey:      privateKey,
		VerificationKey: publicKey,
	}, nil
}

func (c *cryptography) checkKeypairExists(path string) bool {
	exists := false
	if _, err := os.Stat(filepath.Join(path, "private_key.pem")); err == nil {
		exists = true
	}

	if _, err := os.Stat(filepath.Join(path, "public_key.pem")); err == nil {
		exists = true
	}

	return exists
}

func (c *cryptography) generateKeyPair(path string, algorithm Algorithm, keySize int) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0700)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", path, err)
		}
	}

	if c.checkKeypairExists(path) {
		log.Logger.Info("Key pair already exists in " + path)
		return nil
	}

	var err error
	switch algorithm {
	case RS256, RS384, RS512:
		err = c.generateRSAKeyPair(path, keySize)
	case ES256, ES384, ES521:
		err = c.generateECKeyPair(algorithm, path)
	default:
		return errors.New("unsupported algorithm: " + string(algorithm))
	}

	return err
}

func (c *cryptography) generateRSAKeyPair(outputDir string, keySize int) error {
	// Generate RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return ErrGenerateKeyPair
	}

	// Save private key
	privateKeyPath := filepath.Join(outputDir, "private_key.pem")
	privateKeyFile, err := os.Create(privateKeyPath)
	if err != nil {
		return errors.New("failed to create private key file: " + err.Error())
	}

	defer privateKeyFile.Close()

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		return errors.New("failed to write private key to file: " + err.Error())
	}

	// Save public key
	publicKeyPath := filepath.Join(outputDir, "public_key.pem")
	publicKeyFile, err := os.Create(publicKeyPath)
	if err != nil {
		return errors.New("failed to create public key file: " + err.Error())
	}

	defer publicKeyFile.Close()

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return errors.New("failed to marshal public key: " + err.Error())
	}

	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	if err := pem.Encode(publicKeyFile, publicKeyPEM); err != nil {
		return errors.New("failed to write public key to file: " + err.Error())
	}

	return nil
}

func (c *cryptography) generateECKeyPair(algorithm Algorithm, outputDir string) error {
	// Select curve based on algorithm
	var curve elliptic.Curve
	switch algorithm {
	case ES256:
		curve = elliptic.P256()
	case ES384:
		curve = elliptic.P384()
	case ES521:
		curve = elliptic.P521() // Note: P521, not P512
	default:
		log.Logger.Fatal("Unsupported algorithm: " + string(algorithm))
	}

	// Generate EC private key
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return errors.New("failed to generate EC key pair: " + err.Error())
	}

	// Save private key
	privateKeyPath := filepath.Join(outputDir, "private_key.pem")
	privateKeyFile, err := os.Create(privateKeyPath)
	if err != nil {
		return errors.New("failed to create private key file: " + err.Error())
	}

	defer privateKeyFile.Close()

	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return errors.New("failed to marshal private key: " + err.Error())
	}

	privateKeyPEM := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		return errors.New("failed to write private key to file: " + err.Error())
	}

	// Save public key
	publicKeyPath := filepath.Join(outputDir, "public_key.pem")
	publicKeyFile, err := os.Create(publicKeyPath)
	if err != nil {
		return errors.New("failed to create public key file: " + err.Error())
	}

	defer publicKeyFile.Close()

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return errors.New("failed to marshal public key: " + err.Error())
	}

	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	if err := pem.Encode(publicKeyFile, publicKeyPEM); err != nil {
		return errors.New("failed to write public key to file: " + err.Error())
	}

	return nil
}

// LoadPublicKey loads a public key from a file (supports RSA and ECDSA)
func (c *cryptography) LoadPublicKey(path string) (any, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key file: %w", err)
	}

	return c.parsePublicKey(keyData)
}

func (c *cryptography) LoadPrivateKey(path string) (any, error) {
	keyFile, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %w", err)
	}

	return c.parsePrivateKey(keyFile)
}

// ParsePrivateKey parses a private key from PEM-encoded data
func (c *cryptography) parsePrivateKey(keyData []byte) (any, error) {
	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, ErrInvalidKeyFormat
	}

	// Try parsing as PKCS8 (universal format)
	if key, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		switch k := key.(type) {
		case *rsa.PrivateKey, *ecdsa.PrivateKey:
			return k, nil
		default:
			return nil, ErrUnsupportedKey
		}
	}

	// Try parsing as PKCS1 RSA
	if key, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		return key, nil
	}

	// Try parsing as EC
	if key, err := x509.ParseECPrivateKey(block.Bytes); err == nil {
		return key, nil
	}

	return nil, ErrInvalidKeyFormat
}

// ParsePublicKey parses a public key from PEM-encoded data
func (c *cryptography) parsePublicKey(keyData []byte) (any, error) {
	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, ErrInvalidKeyFormat
	}

	// Try parsing as PKIX (universal format)
	if key, err := x509.ParsePKIXPublicKey(block.Bytes); err == nil {
		switch k := key.(type) {
		case *rsa.PublicKey, *ecdsa.PublicKey:
			return k, nil
		default:
			return nil, ErrUnsupportedKey
		}
	}

	return nil, ErrInvalidKeyFormat
}
