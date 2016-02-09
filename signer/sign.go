package signer

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strings"
)

type Signer interface {
	Sign() string
}

type signer struct {
	expire   string
	secret   string
	path     string
	clientIp string
}

func NewSigner(expire, secret, path, clientIp string) Signer {
	return &signer{
		expire:   expire,
		secret:   secret,
		path:     path,
		clientIp: clientIp,
	}
}

func (s *signer) Sign() string {
	str := fmt.Sprintf("%s%s%s %s", s.expire, s.path, s.clientIp, s.secret)
	fmt.Println("str", str)

	//'2147483647/s/link127.0.0.1 secret'

	h := md5.New()

	h.Write([]byte(str))
	base64Str := base64.StdEncoding.EncodeToString(h.Sum(nil))
	finalStr := SanitizeString(base64Str)
	return fmt.Sprintf("%s?md5=%s&expires=%s", s.path, finalStr, s.expire)
}

func SanitizeString(input string) string {
	str := strings.Replace(input, "/", "_", -1)
	str = strings.Replace(str, "+", "-", -1)
	str = strings.Replace(str, "=", "", -1)
	return str
}
