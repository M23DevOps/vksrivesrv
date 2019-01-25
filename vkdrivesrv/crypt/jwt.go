package crypt

import (
	"crypto/hmac"
	"encoding/base64"
	"encoding/json"
	"strings"
	_"fmt"
	"crypto/sha256"
	"errors"
	"time"
	"strconv"
)

type JWTClaim struct {
	SNId	int `json:"sid"`
	SNType	string `json:"stype"`
	Token	string `json:"token"`
}

func JWTSign (claim JWTClaim) (string) {

	json, _ := json.Marshal(claim)

	signingString := EncodeSegment(json)

	hasher := hmac.New(sha256.New, []byte("secret"))
	hasher.Write([]byte(signingString))
	
	signature := EncodeSegment(hasher.Sum(nil))
	return strings.Join([]string{signingString, signature}, ".")
}

func JWTVerify (strSession string) (error) {
	
	parts := strings.Split(strSession, ".")
	if len(parts) != 2 {
		return errors.New("Incorrect session format")
	}

	hasher := hmac.New(sha256.New, []byte("secret"))
	hasher.Write([]byte(parts[0]))

	signature, err := DecodeSegment(parts[1])
	if err != nil {
		return errors.New("Can't Decode Claim")
	}

	if !hmac.Equal(signature, hasher.Sum(nil)) {
		return errors.New("Signature is incorrect")
	}
	return nil
}

func GetSessionToken (intUserId int) (string) {
	hasher := sha256.New()
	hasher.Write([]byte(strconv.Itoa(intUserId) + time.Now().String()))
	return EncodeSegment(hasher.Sum(nil))
}


// Encode JWTClaim specific base64url encoding with padding stripped
func EncodeSegment(seg []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(seg), "=")
}

// Decode JWTClaim specific base64url encoding with padding stripped
func DecodeSegment(seg string) ([]byte, error) {
	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}

	return base64.URLEncoding.DecodeString(seg)
}