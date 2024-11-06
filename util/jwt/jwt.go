package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	env "github.com/Agelessbaby/BloomBlog/util"
	"strings"
	"time"
)

//JWT: Json Web Token

// const (
// 	JWT_SECRET = "0000" //JWT_SECRET has to be stored on the server side properly
// )

var (
	DefautHeader = JwtHeader{
		Algo: "HS256",
		Type: "JWT",
	}
)

type JwtHeader struct {
	Algo string `json:"alg"` // Hash algorithm, default is HMAC SHA256 (written as HS256)
	Type string `json:"typ"` // Token type, standardized as JWT
}

type JwtPayload struct {
	ID          string         `json:"jti"` // JWT ID used to identify this JWT
	Issue       string         `json:"iss"` // Issuer, e.g., WeChat
	Audience    string         `json:"aud"` // Audience, e.g., Honor of Kings
	Subject     string         `json:"sub"` // Subject
	IssueAt     int64          `json:"iat"` // Issued at time, in seconds
	NotBefore   int64          `json:"nbf"` // Not usable before this time, in seconds
	Expiration  int64          `json:"exp"` // Expiration time, in seconds
	UserDefined map[string]any `json:"ud"`  // Other user-defined fields
}

func GenJWT(userId int64) (string, error) {
	header := DefautHeader
	payload := JwtPayload{
		UserDefined: map[string]any{"userid": userId},
		Issue:       "BloomBLog",
		IssueAt:     time.Now().Unix(),
		Expiration:  time.Now().Add(env.TOKEN_EXPIRE).Add(24 * time.Hour).Unix(),
	}
	return genJWT(header, payload, env.JWT_SECRET)
}
func genJWT(header JwtHeader, payload JwtPayload, secret string) (string, error) {
	var part1, part2, signature string
	//header转成json，然后进行Base64编码
	if bs1, err := json.Marshal(header); err != nil {
		return "", err
	} else {
		part1 = base64.RawURLEncoding.EncodeToString(bs1) //这里没有使用StdEncoding，RawURLEncoding的结果中不会包含=+/等url中的特殊字符
	}
	//payload转成json，然后进行Base64编码
	if bs2, err := json.Marshal(payload); err != nil {
		return "", err
	} else {
		part2 = base64.RawURLEncoding.EncodeToString(bs2)
	}
	//基于sha256的哈希认证算法。任意长度的字符串，经过sha256之后长度都变成了256 bits
	h := hmac.New(sha256.New, []byte(secret))
	//signature = HMACSHA256(base64UrlEncode(header) + "." + base64UrlEncode(payload),secret)
	h.Write([]byte(part1 + "." + part2))
	signature = base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	return part1 + "." + part2 + "." + signature, nil
}

func VerifyJwt(token string, secret string) (*JwtHeader, *JwtPayload, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, nil, fmt.Errorf("token has %d parts", len(parts))
	}
	//authenticating hash signature
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(parts[0] + "." + parts[1]))
	signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	if signature != parts[2] { //验证失败
		return nil, nil, fmt.Errorf("authentication failed")
	}

	var part1, part2 []byte
	var err error
	if part1, err = base64.RawURLEncoding.DecodeString(parts[0]); err != nil {
		return nil, nil, fmt.Errorf("header Base64 decoding failed")
	}
	if part2, err = base64.RawURLEncoding.DecodeString(parts[1]); err != nil {
		return nil, nil, fmt.Errorf("payload Base64 decoding failed")
	}

	var header JwtHeader
	var payload JwtPayload
	if err = json.Unmarshal(part1, &header); err != nil {
		return nil, nil, fmt.Errorf("header json decoding failed")
	}
	if err = json.Unmarshal(part2, &payload); err != nil {
		return nil, nil, fmt.Errorf("payload json decoding failed")
	}
	return &header, &payload, nil
}
