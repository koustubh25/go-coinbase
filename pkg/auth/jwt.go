package auth

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math"
	"math/big"
	"os"
	"time"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

type APIKeyClaimsForRESTAPI struct {
	*jwt.Claims
	URI string `json:"uri"`
}

type APIKeyClaimsForWebSocketAPI struct {
	*jwt.Claims
}

var max = big.NewInt(math.MaxInt64)

type nonceSource struct{}

// BuildJWTForRESTAPI is inspired by the code in this document https://docs.cloud.coinbase.com/advanced-trade-api/docs/ws-auth
func BuildJWTForRESTAPI(uri string) (string, error) {
	return buildJWT(uri)
}

func (n nonceSource) Nonce() (string, error) {
	r, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return r.String(), nil
}

func BuildJWTForWebSocketAPI() (string, error) {
	return buildJWT()
}

// buildJWT builds a JWT token for the given URI in case of REST API.
// If no URI is provided, it builds a JWT token for the WebSocket API.
func buildJWT(uri ...string) (string, error) {
	var keyName, keySecret string
	if keyName = os.Getenv("COINBASE_KEY_NAME"); keyName == "" {
		return "", fmt.Errorf("jwt: COINBASE_KEY_NAME is not set")
	}
	if keySecret = os.Getenv("COINBASE_KEY_SECRET"); keySecret == "" {
		return "", fmt.Errorf("jwt: COINBASE_KEY_SECRET is not set")
	}
	block, _ := pem.Decode([]byte(keySecret))
	if block == nil {
		return "", fmt.Errorf("jwt: Could not decode private key")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("jwt: %w", err)
	}

	sig, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.ES256, Key: key},
		(&jose.SignerOptions{NonceSource: nonceSource{}}).WithType("JWT").WithHeader("kid", keyName),
	)
	if err != nil {
		return "", fmt.Errorf("jwt: %w", err)
	}

	jwtClaims := &jwt.Claims{
		Subject:   keyName,
		Issuer:    "coinbase-cloud",
		NotBefore: jwt.NewNumericDate(time.Now()),
		Expiry:    jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
	}
	var cl interface{}
	if len(uri) == 1 {
		cl = &APIKeyClaimsForRESTAPI{
			Claims: jwtClaims,
			URI:    uri[0],
		}
	} else if len(uri) == 0 {
		cl = &APIKeyClaimsForWebSocketAPI{
			Claims: jwtClaims,
		}
	} else {
		return "", fmt.Errorf(`jwt: token can be built only for a single uri`)
	}

	jwtString, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	if err != nil {
		return "", fmt.Errorf("jwt: %w", err)
	}
	return jwtString, nil
}
