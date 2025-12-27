package security

import "errors"

type Algorithm string

const (
	RS256 Algorithm = "RS256"
	RS384 Algorithm = "RS384"
	RS512 Algorithm = "RS512"
	ES256 Algorithm = "ES256"
	ES384 Algorithm = "ES384"
	ES521 Algorithm = "ES521"
	HS256 Algorithm = "HS256"
)

type KeyPair struct {
	SigningKey      any // private key for signing (RSA/ECDSA) or secret (HMAC)
	VerificationKey any // public key for verification (RSA/ECDSA) or secret (HMAC)
}

func AlgorithmFromString(alg string) (Algorithm, error) {
	switch alg {
	case "RS256":
		return RS256, nil
	case "RS384":
		return RS384, nil
	case "RS512":
		return RS512, nil
	case "ES256":
		return ES256, nil
	case "ES384":
		return ES384, nil
	case "ES521":
		return ES521, nil
	case "HS256":
		return HS256, nil
	default:
		return "", errors.New("unsupported algorithm: " + alg)
	}
}
